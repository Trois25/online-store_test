package repository

import (
	"errors"
	"store/features/carts/entity"
	"store/features/carts/model"
	producte "store/features/products/entity"
	productm "store/features/products/model"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type cartRepository struct {
	db      *gorm.DB
	product producte.ProductUseCaseInterface
}

func NewCartRepository(db *gorm.DB, product producte.ProductDataInterface) entity.CartRepositoryInterface {
	return &cartRepository{
		db:      db,
		product: product,
	}
}

// AddProduct implements entity.CartRepositoryInterface.
func (cartRepo *cartRepository) AddCartProduct(cartInput entity.CartsCore, userId string) error {
	//get all product
	var product []productm.Products

	productData, errProduct := cartRepo.product.ReadAllProduct()
	if errProduct != nil {
		return errProduct
	}

	mapProduct := make([]producte.ProductCore, len(productData))
	for i, value := range product {
		mapProduct[i] = producte.ProductCore{
			ID:         value.ID.String(),
			Product:    value.Product,
			Price:      value.Price,
			CategoryId: value.CategoryId,
		}
	}

	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return UUIDerr
	}

	var cartData = model.Carts{
		ID:         newUUID,
		CustomerId: userId,
		ProductId:  cartInput.ProductId,
		Quantity:   cartInput.Quantity,
	}

	//check productId
	checkProductIderr := cartRepo.db.Where("id = ?", cartInput.ProductId).First(&product).Error
	if checkProductIderr != nil {
		return checkProductIderr
	}

	tx := cartRepo.db.Create(&cartData)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// DeleteProduct implements entity.CartRepositoryInterface.
func (cartRepo *cartRepository) DeleteCartProduct(id string, userId string) error {
	var checkData model.Carts

	errData := cartRepo.db.Where("id = ? AND customer_id = ?", id, userId).First(&checkData).Error
	if errData != nil {
		if errors.Is(errData, gorm.ErrRecordNotFound) {
			return errors.New("data not found")
		}
		return errData
	}

	errDelete := cartRepo.db.Delete(&checkData).Error
	if errDelete != nil {
		return errDelete
	}

	return nil
}

// GetAllProduct implements entity.CartRepositoryInterface.
func (cartRepo *cartRepository) GetAllCartProduct(userId string) ([]entity.CartsCore, error) {
	var dataProduct []model.Carts

	errData := cartRepo.db.Where("customer_id = ?", userId).Find(&dataProduct).Error
	if errData != nil {
		return nil, errData
	}

	mapData := make([]entity.CartsCore, len(dataProduct))
	for i, value := range dataProduct {
		mapData[i] = entity.CartsCore{
			ID:         value.ID.String(),
			CustomerId: userId,
			ProductId:  value.ProductId,
			Quantity:   value.Quantity,
			Products:   convertToCartProductCore(value.ProductId, cartRepo.product),
		}
	}

	return mapData, nil
}

// convertToCartProductCore mengambil informasi produk dan mengkonversinya ke dalam bentuk CartProductCore.
func convertToCartProductCore(productId string, productUseCase producte.ProductUseCaseInterface) entity.CartProductCore {
	productInfo, _ := productUseCase.GetProductByID(productId)

	return entity.CartProductCore{
		Product: productInfo.Product,
		Price:   productInfo.Price,
	}
}