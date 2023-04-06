package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) initUsersRoutes(api *gin.RouterGroup) {
	users := api.Group("users", h.userIdentity)
	{
		users.GET("/:userId", h.getUser)
		users.PUT("/:userId", h.updateUser)
		users.DELETE("/:userId", h.deleteUser)
	}
}

func (h *Handler) getUser(ctx *gin.Context) {

	userId, err := getUserId(ctx)
	if err != nil {
		newResponse(ctx, http.StatusBadRequest, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": userId.Hex(),
	})
}

func (h *Handler) updateUser(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}

func (h *Handler) deleteUser(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"message": "good",
	})
}
