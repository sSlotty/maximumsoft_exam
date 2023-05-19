package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"goenv/model"
	"goenv/responses"
	"goenv/services"
	"net/http"
)

type LoginController interface {
	Login() gin.HandlerFunc
}

type loginController struct {
	loginService services.LoginService
	jwtService   services.JWTService
}

func Login() gin.HandlerFunc {
	return func(c *gin.Context) {
		var login model.LoginBody

		if err := c.ShouldBindJSON(&login); err != nil {
			c.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		isUserAuthenticated := services.StaticLoginService().Login(login.Username, login.Password)
		if isUserAuthenticated {
			assessToken := services.JWTAuthService().GenerateToken(login.Username, true, 15)
			refreshToken := services.JWTAuthService().GenerateToken(login.Username, true, 10080)
			c.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Message: "Success Login", Data: map[string]interface{}{"assessToken": assessToken, "refreshToken": refreshToken, "message": "Success Login"}})
		} else {
			c.JSON(http.StatusUnauthorized, responses.ErrorResponse{Status: http.StatusUnauthorized, Message: "Invalid Login"})
		}
	}
}

func RefreshToken() gin.HandlerFunc {
	return func(context *gin.Context) {

		type RefreshToken struct {
			RefreshToken string `json:"refreshToken"`
		}

		var refreshToken RefreshToken
		var body model.LoginBody
		if err := context.ShouldBindJSON(&refreshToken); err != nil {
			context.JSON(http.StatusBadRequest, responses.ErrorResponse{Status: http.StatusBadRequest, Message: err.Error()})
			return
		}

		isRefreshTokenValid, _ := services.JWTAuthService().ValidateToken(refreshToken.RefreshToken)

		if isRefreshTokenValid != nil {
			body.Username = isRefreshTokenValid.Claims.(jwt.MapClaims)["username"].(string)
			newAssessToken := services.JWTAuthService().GenerateToken(body.Username, true, 15)
			newRefreshToken := services.JWTAuthService().GenerateToken(body.Username, true, 10080)
			context.JSON(http.StatusOK, responses.SuccessResponse{Status: http.StatusOK, Message: "Success Refresh Token", Data: map[string]interface{}{"assessToken": newAssessToken, "refreshToken": newRefreshToken, "message": "Success Refresh Token"}})
			return

		}
		context.JSON(http.StatusUnauthorized, responses.ErrorResponse{Status: http.StatusUnauthorized, Message: "Invalid Refresh Token"})
		return

	}
}
