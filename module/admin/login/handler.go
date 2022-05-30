package login

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/xdaawnx/test-privy/helper/constant"
	"github.com/xdaawnx/test-privy/module/admin/user"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
)

// Auth function is for authenticate user
func Auth(ctx echo.Context) error {
	username, password, ok := ctx.Request().BasicAuth()
	layoutFormat := "2006-01-02 15:04:05"

	response := new(constant.Response)

	if !ok {
		response.Status = constant.Authenticationfailed
		response.Responsebuilder()
		return ctx.JSON(http.StatusNotAcceptable, response)
	}
	u, err := user.Auth(username, password)

	if err == nil && u.Email != "" {
		token := jwt.New(jwt.SigningMethodHS256)
		claims := token.Claims.(jwt.MapClaims)
		dt := time.Now().Add(time.Hour * 24)
		claims["iss"] = "test-majoo.id"
		claims["uid"] = fmt.Sprint(u.ID)
		claims["exp"] = fmt.Sprint(dt.Unix())
		claims["username"] = u.Username
		claims["email"] = u.Email
		claims["IssuedAt"] = time.Now().Unix()
		claims["Issuer"] = "backend"
		t, err := token.SignedString([]byte(os.Getenv("JWT_SECRET_KEY")))
		if err == nil {
			success := map[string]interface{}{
				"token_expired_time": dt.Format(layoutFormat),
				"token":              t,
			}
			response.Data = success
			response.Status = constant.Success
			response.Responsebuilder()
			return ctx.JSON(http.StatusOK, response)
		}
	}
	response.Status = constant.Internalerror
	response.Errors = err.Error()
	response.Responsebuilder()
	return ctx.JSON(http.StatusNotFound, response)
}
