package routes

import (
	"net/http"

	"test-privy/module/admin/cake"

	"github.com/labstack/echo/v4"
	m "github.com/labstack/echo/v4/middleware"
)

// Routes is used to declare a route to controller that will be used to handle
func Routes(r *echo.Echo) *echo.Echo {
	apiV1(r)
	return r
}

// apiV1 indicates a versioning function
func apiV1(r *echo.Echo) {
	// cors config to be used in V1
	cors := m.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Authorization"},
		AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodPut},
	}

	// API versioning group with cors in endpoint group
	v := r.Group("/v1/admin", m.CORSWithConfig(cors))
	// v.GET("/auth", login.Auth).Name = "auth.v1"

	// v.Use(echo.WrapMiddleware(auth.JwtValidation))
	v.GET("/cakes", cake.GetAll).Name = "find.cake"
	v.GET("/cakes/:id", cake.Find).Name = "find.detail.cake"
	v.POST("/cakes", cake.Create).Name = "create.cake"
	v.PATCH("/cakes/:id", cake.Update).Name = "update.cake"
	v.DELETE("/cakes/:id", cake.Delete).Name = "delete.cake"

}
