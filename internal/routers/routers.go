// Package routers is a package dedicated to registering routes, and supports both
// manual route registration and automatic route registration.
package routers

import (
	"net/http"

	"github.com/zhufuyi/sponge/docs"
	"github.com/zhufuyi/sponge/internal/config"

	"github.com/zhufuyi/sponge/pkg/gin/handlerfunc"
	"github.com/zhufuyi/sponge/pkg/gin/middleware"
	"github.com/zhufuyi/sponge/pkg/gin/middleware/metrics"
	"github.com/zhufuyi/sponge/pkg/gin/prof"
	"github.com/zhufuyi/sponge/pkg/gin/validator"
	"github.com/zhufuyi/sponge/pkg/jwt"
	"github.com/zhufuyi/sponge/pkg/logger"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var (
	apiV1RouterFns []func(r *gin.RouterGroup) // group router functions
	// if you have other group routes you can define them here
	// example:
	//     apiV2RouterFns []func(r *gin.RouterGroup)
)

// NewRouter create a new router
func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Recovery())
	r.Use(middleware.Cors())

	// request id middleware
	r.Use(middleware.RequestID())

	// logger middleware
	r.Use(middleware.Logging(
		middleware.WithLog(logger.Get()),
		middleware.WithRequestIDFromContext(),
		middleware.WithIgnoreRoutes("/metrics"), // ignore path
	))

	// init jwt middleware
	jwt.Init(
	//jwt.WithExpire(time.Hour*24),
	//jwt.WithSigningKey("123456"),
	//jwt.WithSigningMethod(jwt.SigningMethodHS384),
	)

	// metrics middleware
	if config.Get().App.EnableMetrics {
		r.Use(metrics.Metrics(r,
			//metrics.WithMetricsPath("/metrics"),                // default is /metrics
			metrics.WithIgnoreStatusCodes(http.StatusNotFound), // ignore 404 status codes
		))
	}

	// limit middleware
	if config.Get().App.EnableLimit {
		r.Use(middleware.RateLimit())
	}

	// circuit breaker middleware
	if config.Get().App.EnableCircuitBreaker {
		r.Use(middleware.CircuitBreaker())
	}

	// trace middleware
	if config.Get().App.EnableTrace {
		r.Use(middleware.Tracing(config.Get().App.Name))
	}

	// profile performance analysis
	if config.Get().App.EnableHTTPProfile {
		prof.Register(r, prof.WithIOWaitTime())
	}

	// validator
	binding.Validator = validator.Init()

	r.GET("/health", handlerfunc.CheckHealth)
	r.GET("/ping", handlerfunc.Ping)

	// register swagger routes, generate code via swag init
	docs.SwaggerInfo.BasePath = ""
	// access path /swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// register routers, middleware support
	registerRouters(r, "/api/v1", apiV1RouterFns)
	// if you have other group routes you can add them here
	// example:
	//    registerRouters(r, "/api/v2", apiV2RouteFns, middleware.Auth())

	return r
}

func registerRouters(r *gin.Engine, groupPath string, routerFns []func(*gin.RouterGroup), handlers ...gin.HandlerFunc) {
	rg := r.Group(groupPath, handlers...)
	for _, fn := range routerFns {
		fn(rg)
	}
}
