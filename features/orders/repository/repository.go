package repository

import (
	carte "store/features/carts/entity"
	"store/features/orders/entity"
	"store/features/orders/model"
	producte "store/features/products/entity"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type orderRepository struct {
	db      *gorm.DB
	cart    carte.CartServiceInterface
	product producte.ProductUseCaseInterface
}

func NewOrderRepository(db *gorm.DB, cart carte.CartServiceInterface, product producte.ProductDataInterface) entity.OrderRepositoryInterface {
	return &orderRepository{
		db:      db,
		cart:    cart,
		product: product,
	}
}

// AddOrder implements entity.OrderRepositoryInterface.
func (orderRepo *orderRepository) AddOrder(userId string, cartId string) error {

	// Mendapatkan data keranjang spesifik berdasarkan cartId
	cartData, errCart := orderRepo.cart.GetSpecificCart(userId, cartId)
	if errCart != nil {
		return errCart
	}

	// Mendapatkan semua data produk
	productData, errProduct := orderRepo.product.GetProductByID(cartData.ProductId)
	if errProduct != nil {
		return errProduct
	}

	totalPrice := cartData.Quantity * productData.Price

	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return UUIDerr
	}

	// Membuat data pesanan baru dengan total harga yang telah dihitung
	var orderData = model.Orders{
		ID:         newUUID,
		UserId:     userId,
		CartId:     cartId,
		TotalPrice: totalPrice,
	}

	
	txDel := orderRepo.cart.DeleteCartProduct(cartId, userId)
	if txDel != nil {
		return txDel
	}

	// Menyimpan data pesanan ke dalam database
	tx := orderRepo.db.Create(&orderData)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

