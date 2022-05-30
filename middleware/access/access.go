package access

// Routes is a variable to save a route on this project

import (
	"github.com/labstack/echo/v4"
)

// Routes is a variable to save a route on this project
var Routes []*echo.Route

// CanAccess is used to find the user access can
type CanAccess struct {
	Method  string
	Path    string
	Type    string
	Create  bool
	Find    bool
	Update  bool
	Delete  bool
	Approve bool
}

// GenerateRoutes is used to save a route data to the variable
func GenerateRoutes(e *echo.Echo) echo.MiddlewareFunc {
	for _, y := range e.Routes() {
		if y.Name != "github.com/labstack/echo/v4.glob..func1" {
			Routes = append(Routes, y)
		}
	}

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return next
	}
}
