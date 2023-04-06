package v1

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/korasdor/go-ether-test/pkg/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	authorizationHeader = "Authorization"

	userCtx = "userId"
)

func (h *Handler) userIdentity(ctx *gin.Context) {
	td, err := h.parseAuthHeader(ctx)
	if err != nil {
		newResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	tb := &auth.TokenBinding{}
	tb, err = tb.Parse(ctx.Request)
	if err != nil {
		newResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	if td.TokenBinding.IPAddr != tb.IPAddr || td.TokenBinding.UserAgent != tb.UserAgent {
		newResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, td.UserId)
}

func (h *Handler) parseAuthHeader(ctx *gin.Context) (*auth.TokenData, error) {
	header := ctx.GetHeader(authorizationHeader)
	if header == "" {
		return nil, errors.New("empty auth header")
	}

	headerParts := strings.Split(header, " ")
	if len(headerParts) != 2 || headerParts[0] != "Bearer" {
		return nil, errors.New("invalid auth header")
	}

	if len(headerParts[1]) == 0 {
		return nil, errors.New("token is empty")
	}

	return h.tokenManager.ParseJWT(headerParts[1])
}

func getUserId(c *gin.Context) (primitive.ObjectID, error) {
	return getIdByContext(c, userCtx)
}

func getIdByContext(c *gin.Context, context string) (primitive.ObjectID, error) {
	idFromCtx, ok := c.Get(context)
	if !ok {
		return primitive.ObjectID{}, errors.New("studentCtx not found")
	}

	idStr, ok := idFromCtx.(string)
	if !ok {
		return primitive.ObjectID{}, errors.New("studentCtx is of invalid type")
	}

	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return id, nil
}
