package middlewares

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type GoogleKey struct {
	Keys []struct {
		Kid string `json:"kid"`
		Kty string `json:"kty"`
		Alg string `json:"alg"`
		Use string `json:"use"`
		N   string `json:"n"`
		E   string `json:"e"`
	} `json:"keys"`
}

func DownloadGoogleKey() (*GoogleKey, error) {
	resp, err := http.Get(GOOGLE_KEY_URL)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var googleKey GoogleKey
	err = json.NewDecoder(resp.Body).Decode(&googleKey)
	if err != nil {
		return nil, err
	}

	return &googleKey, nil
}

func JwtGoogle() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.Request.Header.Get(AUTHORIZATION)
		if authorizationHeader == "" {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
		}
		tokenParts := strings.Split(authorizationHeader, "Bearer ")
		if len(tokenParts) != 2 {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
		}
		idToken := tokenParts[1]
		googleKey, err := DownloadGoogleKey()
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
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
		claims := token.Claims.(jwt.MapClaims)
		emailVerified := claims["email_verified"].(bool)
		if !emailVerified {
			ctx.JSON(http.StatusUnauthorized, nil)
			ctx.Abort()
		}
		ctx.Next()
	}
}
