package service

import (
	"errors"
	"store/features/carts/entity"
)

type cartService struct {
	cartRepository entity.CartRepositoryInterface
}

func NewCartService(cart entity.CartRepositoryInterface) entity.CartServiceInterface {
	return &cartService{
		cartRepository: cart,
	}
}

// AddProduct implements entity.CartServiceInterface.
func (cart *cartService) AddCartProduct(cartInput entity.CartsCore, userId string) error {
	if userId == "" {
		return errors.New("you need to login first")
	}

	if cartInput.ProductId == "" {
		return errors.New("product id can't empty")
	}

	if cartInput.Quantity < 0 {
		return errors.New("product must be more than 0")
	}

	errAdd := cart.cartRepository.AddCartProduct(cartInput, userId)
	if errAdd != nil {
		return errAdd
	}

	return nil
}

// DeleteProduct implements entity.CartServiceInterface.
func (cart *cartService) DeleteCartProduct(id string, userId string) error {
	if id == "" {
		return errors.New("cart id can't empty")
	}

	errPurchase := cart.cartRepository.DeleteCartProduct(id, userId)
	if errPurchase != nil {
		return errors.New("can't delete product")
	}

	return nil
}

// GetAllProduct implements entity.CartServiceInterface.
func (cart *cartService) GetAllCartProduct(userId string) ([]entity.CartsCore, error) {
	products, err := cart.cartRepository.GetAllCartProduct(userId)
	if err != nil {
		return nil, errors.New("error get data")
	}
	return products, nil
}

// GetSpecificCart implements entity.CartServiceInterface.
func (cart *cartService) GetSpecificCart(userId string, id string) (entity.CartsCore, error) {
	if id == "" {
		return entity.CartsCore{}, errors.New("id invalid")
	}

	if userId == "" {
		return entity.CartsCore{}, errors.New("you need to login first")
	}

	cartData, err := cart.cartRepository.GetSpecificCart(userId, id)
	if err != nil {
		return entity.CartsCore{}, errors.New("can't read data")
	}

	return cartData, nil
}
