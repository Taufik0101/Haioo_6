package controller

import (
	"Haioo_6/api/dto"
	"Haioo_6/api/service"
	"Haioo_6/api/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type CartController interface {
	GetCart(ctx *gin.Context)
	CreateCart(ctx *gin.Context)
	DeleteCart(ctx *gin.Context)
}

type cartController struct {
	cartService service.CartService
}

func (c cartController) GetCart(ctx *gin.Context) {
	//TODO implement me
	var DTOCart dto.ListCartInput
	_ = ctx.ShouldBind(&DTOCart)

	var carts []*dto.ListCartResponse
	carts, _ = c.cartService.GetCart(DTOCart)
	resp := utils.BuildResponse(true, "Get Data Cart Berhasil", carts)
	ctx.JSON(http.StatusOK, resp)
}

func (c cartController) CreateCart(ctx *gin.Context) {
	//TODO implement me
	var DTOCart dto.CreateCart
	errCreate := ctx.ShouldBind(&DTOCart)
	if errCreate != nil {
		response := utils.BuildErrorResponse("Failed to parsing", errCreate.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
	} else {
		res, err := c.cartService.CreateCart(DTOCart)
		if err != nil {
			response := utils.BuildErrorResponse("Tambah Data Gagal", err.Error(), utils.EmptyObj{})
			ctx.JSON(http.StatusBadRequest, response)
			return
		}
		resp := utils.BuildResponse(true, "Tambah Data Berhasil", res)
		ctx.JSON(http.StatusOK, resp)
	}
}

func (c cartController) DeleteCart(ctx *gin.Context) {
	//TODO implement me
	id := ctx.Query("kodeProduk")
	res, err := c.cartService.DeleteCart(id)
	if !res || err != nil {
		response := utils.BuildErrorResponse("Failed to delete", err.Error(), utils.EmptyObj{})
		ctx.JSON(http.StatusBadRequest, response)
		return
	}
	resp := utils.BuildResponse(true, "Delete Data Berhasil", res)
	ctx.JSON(http.StatusOK, resp)
}

func NewCartController(cartServ service.CartService) CartController {
	return &cartController{cartService: cartServ}
}
