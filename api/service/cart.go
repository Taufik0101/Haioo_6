package service

import (
	"Haioo_6/api/dto"
	"Haioo_6/api/model"
	"Haioo_6/api/utils"
	"fmt"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type CartService interface {
	GetCart(input dto.ListCartInput) ([]*dto.ListCartResponse, error)
	CreateCart(input dto.CreateCart) (*model.Cart, error)
	DeleteCart(id string) (bool, error)
}

type cartService struct {
	cartConnection *gorm.DB
}

func (c cartService) GetCart(input dto.ListCartInput) ([]*dto.ListCartResponse, error) {
	var cartsResponse []*dto.ListCartResponse
	var carts []*model.Cart

	queryCart := c.cartConnection.Model(&model.Cart{})

	queryCart.Scopes(func(db *gorm.DB) *gorm.DB {
		if input.Filter != nil {
			if input.Filter.NamaProduk != nil && utils.DerefString(input.Filter.NamaProduk) != "" {
				db.Where(model.Cart{NameProduct: utils.DerefString(input.Filter.NamaProduk)})
			}

			if input.Filter.Kuantitas != nil && *input.Filter.Kuantitas != 0 {
				db.Where(model.Cart{Quantity: *input.Filter.Kuantitas})
			}
		}

		return db
	})

	queryCart.Order(clause.OrderByColumn{
		Column: clause.Column{
			Name: "created_at",
		},
		Desc: true,
	}).Find(&carts)

	for _, val := range carts {
		cartsResponse = append(cartsResponse, &dto.ListCartResponse{
			Value: fmt.Sprintf("%s - %s - (%d)", val.KodeProduct, val.NameProduct, val.Quantity),
		})
	}

	return cartsResponse, nil
}

func (c cartService) CreateCart(input dto.CreateCart) (*model.Cart, error) {
	//TODO implement me
	var cart *model.Cart

	queryCart := c.cartConnection.Model(&model.Cart{})

	err := queryCart.Where(model.Cart{KodeProduct: input.KodeProduk}).First(&cart).Error

	if err != nil {
		newCart := &model.Cart{
			KodeProduct: input.KodeProduk,
			NameProduct: input.NamaProduk,
			Quantity:    input.Kuantitas,
		}

		errSave := c.cartConnection.Create(newCart).Error

		if errSave != nil {
			return nil, errSave
		}

		return newCart, nil
	} else {
		updateCart := &model.Cart{
			Quantity: cart.Quantity + input.Kuantitas,
		}

		errUpdate := queryCart.Where(model.Cart{ID: cart.ID}).Updates(updateCart).Error

		if errUpdate != nil {
			return nil, errUpdate
		}

		queryCart.Where(model.Cart{KodeProduct: input.KodeProduk}).First(&cart)

		return cart, nil
	}
}

func (c cartService) DeleteCart(id string) (bool, error) {
	//TODO implement me
	var cart *model.Cart

	queryCart := c.cartConnection.Model(&model.Cart{})

	errDelete := queryCart.Where(model.Cart{KodeProduct: id}).Delete(&cart).Error

	if errDelete != nil {
		return false, errDelete
	}

	return true, nil
}

func NewCartService(cartConn *gorm.DB) CartService {
	return &cartService{cartConnection: cartConn}
}
