package rest

import (
	"fmt"
	"net/http"

	"github.com/Tonmoy404/Smart-Inventory/logger"
	"github.com/Tonmoy404/Smart-Inventory/util"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const (
	authorizationHeaderKey  = "token"
	authorizationPayloadKey = "authorization_payload"
)

func (s *Server) authMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie(authorizationHeaderKey)
		if err != nil {
			logger.Error(ctx, "error in getting cookie", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, s.svc.Error(ctx, util.EN_UNAUTHENTICATED_ERROR, "Unauthorized"))
			return
		}

		if len(accessToken) == 0 {
			logger.Error(ctx, "authorization header is not provided", nil)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, s.svc.Error(ctx, util.EN_UNAUTHENTICATED_ERROR, "Unauthorized"))
			return
		}

		// parse the JWT.
		token, err := jwt.Parse(accessToken, s.validateJWT)
		if err != nil {
			logger.Error(ctx, "failed to parse the JWT", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, s.svc.Error(ctx, util.EN_UNAUTHENTICATED_ERROR, "Unauthorized"))
			return
		}

		// Check if the token is valid.
		if !token.Valid {
			logger.Error(ctx, "failed to create JWKS from resource at the given URL", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, s.svc.Error(ctx, util.EN_UNAUTHENTICATED_ERROR, "Unauthorized"))
			return
		}

		logger.Info(ctx, "The token is valid", nil)

		// Get the token claims.
		claims := token.Claims.(jwt.MapClaims)

		payload := Payload{
			ID:          claims["id"].(string),
			UserName:    claims["username"].(string),
			Email:       claims["email"].(string),
			Role:        claims["role"].(string),
			FullName:    claims["full_name"].(string),
			PhoneNumber: claims["phone_number"].(string),
			Status:      claims["status"].(string),
		}

		logger.Info(ctx, "User Payload", payload)

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

func (s *Server) validateJWT(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
	}

	return []byte(s.jwt.SecretKey), nil
}

func corsMiddleware(c *gin.Context) {
	// Get the origin from the request header
	origin := c.Request.Header.Get("Origin")

	// Set the Access-Control-Allow-Origin header to the origin of the client application
	c.Writer.Header().Set("Access-Control-Allow-Origin", origin)

	// Allow specific headers
	c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization, token")

	// Allow all methods
	c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, PATCH, DELETE, OPTIONS")

	// Allow credentials
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")

	// Handle OPTIONS method
	if c.Request.Method == "OPTIONS" {
		c.AbortWithStatus(204)
		return
	}

	c.Next()
}

func (s *Server) isAuthorizedMiddleware(action, object string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		logger.Info(ctx, "action and object", fmt.Sprintf("%v %v", action, object))
		// decode cookie
		authPayload := ctx.MustGet(authorizationPayloadKey).(Payload)

		// check permission
		permitted := s.svc.IsPermitted(ctx, authPayload.Role, action, object)

		if !permitted {
			logger.Error(ctx, "Forbidden for user", nil)
			ctx.AbortWithStatusJSON(http.StatusForbidden, s.svc.Error(ctx, util.EN_FORBIDDEN_ERROR, "Forbidden"))
			return
		}

		ctx.Next()
	}
}

//set payload as nil if user is public else set valid payload for system user
func (s *Server) isSystemUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		accessToken, err := ctx.Cookie(authorizationHeaderKey)
		if err != nil || len(accessToken) == 0 {
			ctx.Set(authorizationPayloadKey, nil)
			ctx.Next()
		}

		// parse the JWT.
		token, err := jwt.Parse(accessToken, s.validateJWT)
		if err != nil {
			logger.Error(ctx, "failed to parse the JWT", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, s.svc.Error(ctx, util.EN_UNAUTHENTICATED_ERROR, "Unauthorized"))
			return
		}

		// Check if the token is valid.
		if !token.Valid {
			logger.Error(ctx, "failed to create JWKS from resource at the given URL", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, s.svc.Error(ctx, util.EN_UNAUTHENTICATED_ERROR, "Unauthorized"))
			return
		}

		logger.Info(ctx, "The token is valid", nil)

		// Get the token claims.
		claims := token.Claims.(jwt.MapClaims)

		payload := Payload{
			ID:          claims["id"].(string),
			UserName:    claims["username"].(string),
			Email:       claims["email"].(string),
			Role:        claims["role"].(string),
			FullName:    claims["full_name"].(string),
			PhoneNumber: claims["phone_number"].(string),
			Status:      claims["status"].(string),
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}
