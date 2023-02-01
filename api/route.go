package api

import (
	"Haioo_6/api/controller"
	"github.com/gin-gonic/gin"
)

type Route interface {
	Routes(r *gin.Engine)
}

type route struct {
	cartController controller.CartController
}

func (r2 route) Routes(r *gin.Engine) {
	//TODO implement me
	cart := r.Group("/cart")
	{
		cart.POST("/", r2.cartController.GetCart)
		cart.POST("/add", r2.cartController.CreateCart)
		cart.DELETE("/delete", r2.cartController.DeleteCart)
	}
}

func NewRoute(cartCont controller.CartController) Route {
	return &route{cartController: cartCont}
}
