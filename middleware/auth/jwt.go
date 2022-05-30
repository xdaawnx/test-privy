package auth

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/xdaawnx/test-privy/helper/constant"

	"github.com/johansetia/jowt"
)

type er map[string]interface{}

// JwtValidation func is used to verify a token that has been gives in login func
func JwtValidation(next http.Handler) http.Handler {

	response := new(constant.Response)

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authorizationHeader := r.Header.Get("Authorization")
		if !strings.Contains(authorizationHeader, "Bearer") {
			response.Status = constant.Authenticationfailed
			response.Responsebuilder()
			stop(w, response)
		} else {
			tokenString := strings.Replace(authorizationHeader, "Bearer ", "", -1)
			verify := jowt.Verify(os.Getenv("JWT_SECRET_KEY")).SetToken(tokenString)

			if verify.Status() {
				if expired(verify.Payload) {
					response.Status = constant.Dataexpired
					response.Responsebuilder()
					stop(w, response)
				} else {
					r.Header.Set("uid", fmt.Sprint(verify.Payload["uid"]))
					next.ServeHTTP(w, r)
				}
			} else {
				response.Status = constant.Authenticationfailed
				response.Responsebuilder()
				stop(w, response)
			}
		}
	})
}

func expired(p jowt.Payload) bool {
	now := time.Now().Unix()
	expired, err := strconv.ParseInt(p["exp"].(string), 10, 64)
	if err != nil {
		return true
	}
	if now > expired {
		return true
	}
	return false
}

func stop(w http.ResponseWriter, res *constant.Response) {
	w.Header().Set("Content-Type", "application/json")
	m, err := json.Marshal(res)
	errStr := string(m)
	if err != nil {
		errStr = `{"error":true}`
	}
	http.Error(w, errStr, http.StatusNotFound)
	return
}
