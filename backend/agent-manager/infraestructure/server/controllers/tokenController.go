package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/LuisDiazM/agent-manager/domain/usecases/userUsecase/entities"
	"github.com/LuisDiazM/agent-manager/infraestructure/app"
	"github.com/LuisDiazM/agent-manager/infraestructure/server/middlewares"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func decodeToken(idToken string) *jwt.Token {
	googleKey, err := middlewares.DownloadGoogleKey()
	if err != nil {
		return nil
	}
	token, _ := jwt.Parse(idToken, func(token *jwt.Token) (interface{}, error) {
		kid := token.Header["kid"].(string)
		for _, key := range googleKey.Keys {
			if key.Kid == kid {
				return jwt.ParseRSAPublicKeyFromPEM([]byte(fmt.Sprintf("-----BEGIN CERTIFICATE-----\n%s\n-----END CERTIFICATE-----", key.N)))
			}
		}
		return nil, fmt.Errorf("public key not found")
	})
	return token
}

func LoginController(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToken := ctx.Query("id")
		token := decodeToken(idToken)
		if token == nil {
			ctx.JSON(http.StatusNoContent, nil)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		id := claims["email"].(string)
		user := app.UserUsecase.GetUserById(id, ctx)
		if user == nil {
			ctx.JSON(http.StatusNoContent, nil)
			return
		} else {
			ctx.JSON(http.StatusOK, user)
		}
	}
}

func RegisterController(app *app.Application) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		idToken := ctx.Query("id")
		token := decodeToken(idToken)
		if token == nil {
			ctx.JSON(http.StatusBadRequest, nil)
			return
		}
		claims := token.Claims.(jwt.MapClaims)
		email := claims["email"].(string)
		name := claims["name"].(string)
		picture := claims["picture"].(string)
		var user entities.User = entities.User{Name: name, Email: email, Id: email, Picture: picture, CreatedAt: time.Now().UTC()}
		response := app.UserUsecase.InsertUser(user, ctx)
		if response != nil {
			reference := *response
			if reference["exists_previous_licenses"] {
				ctx.JSON(http.StatusOK, gin.H{"user_id": nil, "is_created": false})
				return
			}
			ctx.JSON(http.StatusCreated, gin.H{"user_id": user.Id, "is_created": true})

		} else {
			ctx.JSON(http.StatusOK, gin.H{"user_id": nil, "is_created": false})
		}

	}
}
