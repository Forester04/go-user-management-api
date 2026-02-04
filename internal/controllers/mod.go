package controllers

import (
	"net/http"
	"reflect"
	"strings"

	"github.com/Forester04/go-user-management-api/internal/services"
	"github.com/Forester04/go-user-management-api/internal/viewmodel"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator"
	"go.uber.org/zap"
)

type Router struct {
	engine *gin.Engine
	logger *zap.Logger
}

func NewRouter(logger *zap.Logger, svc services.ServiceInterface) *Router {
	router := &Router{}

	router.logger = logger
	router.engine = gin.Default()
	config(router.engine)

	/* Middleswares */
	router.engine.Use(router.corsMiddleware())
	router.engine.Use(router.responseViewmodelMiddleware())
	router.engine.Use(router.errorHandlerMiddleware())

	router.registerRoutes(svc)

	return router
}

func (rtr *Router) Run(addr string) error {
	return rtr.engine.Run(addr)
}

func (rtr *Router) registerRoutes(svc services.ServiceInterface) {
	/* Auth */
	auth := rtr.engine.Group("/auth")
	registerAuthRoutes(auth, svc)

	/* User */
	user := rtr.engine.Group("/user")
	registerUserRoutes(user, svc)

	rtr.engine.GET("/health", healthController())
}

func healthController() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response := &viewmodel.HealthResponse{}

		response.Body.Message = "health OK"

		ctx.Set(ContextKeyStatusCode, http.StatusOK)
		ctx.Set(ContextKeyResponseviewmodel, response)
	}
}

func config(router *gin.Engine) {
	router.RedirectTrailingSlash = false

	// Custom validator
	// This is used to be able to get the tag name instead of the struct field name
	// to send it to thw client, and the client can use it to display the error message
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return fld.Name
			}
			return name
		})
	}
}
