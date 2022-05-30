package customcfg

import (
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	em "github.com/labstack/echo/v4/middleware"
)

// DefaultErr is used to mask an error by echo
func DefaultErr() func(err error, c echo.Context) {
	return func(err error, c echo.Context) {
		report, ok := err.(*echo.HTTPError)
		if !ok {
			report = echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}
		cerr := map[string]interface{}{
			"error":   true,
			"message": report.Message,
		}
		c.JSON(report.Code, cerr)
	}
}

// LoggerCfg is used to show a log access
func LoggerCfg() echo.MiddlewareFunc {
	return em.LoggerWithConfig(
		em.LoggerConfig{
			Format:           "${time_custom} | ${status}  ⇨ ${method} ${protocol} ${host}${uri} ${latency_human} \n",
			CustomTimeFormat: "15:04:05",
		})
}

// RecoverCfg is used to show a log access
func RecoverCfg() echo.MiddlewareFunc {
	var disStack bool
	var disPrint bool

	if os.Getenv("ENVIRONMENT") != "DEV" {
		disStack = true
		disPrint = true
	}

	return em.RecoverWithConfig(em.RecoverConfig{
		StackSize:         5 << 10,
		DisableStackAll:   disStack,
		DisablePrintStack: disPrint,
	})
}

// Secure is used to securing from another content
func Secure() echo.MiddlewareFunc {
	return em.SecureWithConfig(em.SecureConfig{
		XSSProtection:         "1; mode=block",
		ContentTypeNosniff:    "nosniff",
		ContentSecurityPolicy: "test-privy.id",
	})
}
