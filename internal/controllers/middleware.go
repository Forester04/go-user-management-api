package controllers

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"strconv"

	"github.com/Forester04/go-user-management-api/internal/errcode"
	"github.com/Forester04/go-user-management-api/internal/viewmodel"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

func requestViewmodelMiddleware(requestViewmodel interface{}) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Get the body
		requestBody, err := io.ReadAll(ctx.Request.Body)
		if err != nil {
			ctx.Error(fmt.Errorf("%w: %v", errcode.ErrInvalidParameters, err))
			ctx.Abort()
			return
		}
		// transform body
		transformedBody := []byte(`{"body":` + string(requestBody) + `}`)
		ctx.Request.Body = io.NopCloser(bytes.NewReader(transformedBody))

		requestViewmodelInstance := reflect.New(reflect.TypeOf(requestViewmodel).Elem()).Interface()

		err = bindURITaggedFields(ctx, requestViewmodelInstance)
		if err != nil {
			ctx.Error(err)
			ctx.Abort()
			return
		}

		// binding the requestBody, query, headers
		if err := ctx.ShouldBind(requestViewmodelInstance); err != nil {
			//Rewrite the original body
			ctx.Request.Body = io.NopCloser(bytes.NewReader(requestBody))

			var verr validator.ValidationErrors
			if errors.As(err, &verr) {
				failedFields := make(map[string]string, len(verr))
				for _, fieldError := range verr {
					failedFields[fieldError.Field()] = fieldError.Namespace()
				}
				ctx.Set(ContextKeyInvalidFields, failedFields)
			}
			err = fmt.Errorf("%w: %v", errcode.ErrInvalidParameters, err)
			ctx.Error(err)
			ctx.Abort()
			return
		}

		ctx.Request.Body = io.NopCloser(bytes.NewReader(requestBody))

		ctx.Set(ContextKeyRequestViewmodel, requestViewmodelInstance)
		ctx.Next()
	}
}

func bindURITaggedFields(ctx *gin.Context, data interface{}) error {
	val := reflect.ValueOf(data).Elem()
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		tag := typ.Field(i).Tag.Get("uri")

		if tag == "" {
			continue
		}

		paramValue, exists := ctx.Params.Get(tag)
		if !exists {
			continue
		}

		switch field.Kind() {
		case reflect.String:
			field.SetString(paramValue)
		case reflect.Uint:
			paramUint, err := strconv.ParseUint(paramValue, 10, 32)
			if err != nil {
				return errcode.ErrInvalidParameters
			}
			field.SetUint(paramUint)
		case reflect.Int:
			paramInt, err := strconv.Atoi(paramValue)
			if err != nil {
				return errcode.ErrInvalidParameters
			}
			field.SetInt(int64(paramInt))
		default:
			continue
		}

	}
	return nil
}

func (rtr *Router) responseViewmodelMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		if responseViewmodel, exists := ctx.Get(ContextKeyResponseviewmodel); exists {
			statusCode := ctx.GetInt(ContextKeyStatusCode)

			value := reflect.ValueOf(responseViewmodel)
			if value.Kind() == reflect.Ptr && !value.IsNil() {
				bodyField := value.Elem().FieldByName("Body")
				if bodyField.IsValid() {
					ctx.JSON(statusCode, bodyField.Interface())
					return
				}
				ctx.Status(statusCode)
			}
		}
		ctx.Status(http.StatusBadRequest)
	}
}

func (rtr *Router) errorHandlerMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()

		err := ctx.Errors.Last()
		if err != nil {
			body, errReading := io.ReadAll(ctx.Request.Body)
			if errReading != nil {
				body = []byte("error reading body")
			}

			rtr.logger.Error("middleware error",
				zap.Error(err.Err),
				zap.String("path", ctx.Request.URL.Path),
				zap.String("method", ctx.Request.Method),
				zap.String("ip", ctx.ClientIP()),
				zap.String("body", string(body)))

			var GoCleanError errcode.GoCleanError
			if errors.As(err.Err, &GoCleanError) {
				response := &viewmodel.BadRequestErrorResponse{}
				response.Body.Message = GoCleanError.Error()

				//check if error is invalid parameter
				if errors.Is(GoCleanError, errcode.ErrInvalidParameters) {
					failedFields := ctx.GetStringMapString(ContextKeyInvalidFields)
					if len(failedFields) != 0 {
						response.Body.Context = failedFields
					}
				}
				ctx.Set(ContextKeyStatusCode, http.StatusBadRequest)
				ctx.Set(ContextKeyResponseviewmodel, response)
				return
			}
			// if the error is not a GoClean, return generic error
			response := &viewmodel.InternalServerErrorResponse{}
			response.Body.Message = "internal error"
			ctx.Set(ContextKeyStatusCode, http.StatusInternalServerError)
			ctx.Set(ContextKeyResponseviewmodel, response)
			return
		}
	}
}

// Set the CORS rules
func (rtr *Router) corsMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		ctx.Writer.Header().Set("Access-Control-Max-Age", "86400")
		ctx.Writer.Header().
			Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE, PATCH")
		ctx.Writer.Header().
			Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, Accept, Authorization, Two-Factor-Code, Recaptcha, Lang, Country, Session-Id, Api-Key")
		ctx.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		ctx.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(200)
		} else {
			ctx.Next()
		}
	}
}
