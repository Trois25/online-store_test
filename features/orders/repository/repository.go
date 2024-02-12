package repository

import (
	"errors"
	carte "store/features/carts/entity"
	"store/features/orders/entity"
	"store/features/orders/model"
	producte "store/features/products/entity"
	"time"

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

// GetSpecificOrder implements entity.OrderRepositoryInterface.
func (orderRepo *orderRepository) GetSpecificOrder(userId string, orderId string) (entity.OrdersCore, error) {
	var order model.Orders

	err := orderRepo.db.Where("id = ?", orderId).First(&order).Error
	if err != nil {
		return entity.OrdersCore{}, err
	}

	orderCore := entity.OrdersCore{
		ID:         order.ID.String(),
		UserId:     userId,
		CartId:     order.CartId,
		TotalPrice: order.TotalPrice,
	}

	return orderCore, nil
}

// EditOrder implements entity.OrderRepositoryInterface.
func (orderRepo *orderRepository) EditOrder(userId string, orderId string) error {

	var orderData model.Orders
	errData := orderRepo.db.Where("id = ?", orderId).First(&orderData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return errors.New("record not found")
		}
		return errData
	}

	uuidID, err := uuid.Parse(orderId)
	if err != nil {
		return err
	}

	orderData.ID = uuidID
	orderData.PaymentStatus = "Paid"
	orderData.UpdatedAt = time.Now()

	var update = model.Orders{
		ID:            orderData.ID,
		UserId:        orderData.UserId,
		CartId:        orderData.CartId,
		PaymentStatus: orderData.PaymentStatus,
		TotalPrice:    orderData.TotalPrice,
		UpdatedAt:     orderData.UpdatedAt,
		CreatedAt: orderData.CreatedAt,
	}

	errSave := orderRepo.db.Save(&update)
	if errData != nil {
		return errSave.Error
	}

	return nil
}
