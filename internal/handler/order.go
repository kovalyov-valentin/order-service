package handler

import (
	"context"
	"errors"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kovalyov-valentin/orders-service/pkg/lib/api/response"
	"github.com/kovalyov-valentin/orders-service/internal/models"
)

func (h *Handler) Create(ctx *gin.Context) {

	var order models.Order
	if err := ctx.Bind(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, response.Error{Error: err.Error()})
		return
	}

	if err := h.services.Repo.Create(context.TODO(), order.OrderUID, order); err != nil {
		ctx.JSON(http.StatusInternalServerError, response.Error{Error: err.Error()})
		return

	}

	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, map[string]interface{}{
		"uid": order.OrderUID,
	})

}

func (h *Handler) GetOrderByUID(ctx *gin.Context)  {
	uid := ctx.Param("orderuid")

	if uid == "" {
		ctx.JSON(http.StatusInternalServerError, errors.New("чтобы получить order укажите его uid").Error())
		return
	}
	order, err := h.services.Repo.GetOrderByUID(context.TODO(), uid)
	if err != nil {
		log.Println("failed to get order")
		response.NewErrorResponse(ctx, http.StatusBadRequest, "error: no order with your uid")
		return
	}

	ctx.Writer.Header().Add("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, order)
}
