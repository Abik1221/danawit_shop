package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("error can't find product")
	ErrCantDecodeProducts = errors.New("error can't decode products")
	ErrUserIdIsNotValid   = errors.New("error user Id is not valid")
	ErrCantUpdateUser     = errors.New("error Can't Update user")
	ErrCantRemoveItemCart = errors.New("error can't remove item cart")
	ErrCantGetItem        = errors.New("error can't get item")
	ErrCantBuyItem        = errors.New("error can't buy item")
)

func AddToCart() {

}

func RemoveFromCart() {

}

func BuyItemFromCart() {

}

func InstantBuy() {

}
