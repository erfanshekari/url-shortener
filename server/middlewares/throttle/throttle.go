package throttle

import (
	"net/http"

	"github.com/erfanshekari/url-shortener/config"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/time/rate"
)

func GetThrottleMiddleware() echo.MiddlewareFunc {

	throttleConfig := config.GetConfigInstance().Server.Throttle

	var ThrottleConfig middleware.RateLimiterConfig = middleware.RateLimiterConfig{
		Skipper: middleware.DefaultSkipper,
		Store: middleware.NewRateLimiterMemoryStoreWithConfig(
			middleware.RateLimiterMemoryStoreConfig{
				Rate:      rate.Limit(throttleConfig.Rate),
				Burst:     throttleConfig.Burst,
				ExpiresIn: throttleConfig.ExpiresIn,
			},
		),
		IdentifierExtractor: func(ctx echo.Context) (string, error) {
			id := ctx.RealIP()
			return id, nil
		},
		ErrorHandler: func(context echo.Context, err error) error {
			return context.JSON(http.StatusForbidden, nil)
		},
		DenyHandler: func(context echo.Context, identifier string, err error) error {
			return context.JSON(http.StatusTooManyRequests, nil)
		},
	}

	return middleware.RateLimiterWithConfig(ThrottleConfig)
}
