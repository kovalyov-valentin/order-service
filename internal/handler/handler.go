package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/kovalyov-valentin/orders-service/internal/service"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.Default()

	router.LoadHTMLFiles("./web/static/index.html")

	orders := router.Group("/orders")
	{
		orders.GET("", func(ctx *gin.Context) {
			ctx.HTML(200, "index.html", map[string]string{"title": "home page"})
		})
		orders.POST("", h.Create)
		orders.GET(":orderuid", h.GetOrderByUID)
	}

	return router
}

