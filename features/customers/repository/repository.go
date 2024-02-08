package repository

import (
	"errors"
	customers "store/features/customers/entity"
	"store/features/customers/model"
	crypt "store/utils/bcrypt"
	middlewares "store/utils/jwt"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func New(db *gorm.DB) customers.CustomerDataInterface {
	return &customerRepository{
		db: db,
	}
}

// Login implements entity.CustomerDataInterface.
func (customerRep *customerRepository) Login(email string, password string) (customers.CustomerCore, string, error) {

	var data model.Customers

	tx := customerRep.db.Where("email = ? ", email).First(&data)
	if tx.Error != nil {
		return customers.CustomerCore{}, "", tx.Error
	}

	var token string

	if tx.RowsAffected > 0 {
		if crypt.CheckPasswordHash(data.Password, password) {
			var errToken error
			token, errToken = middlewares.CreateToken(data.ID)
			if errToken != nil {
				return customers.CustomerCore{}, "", errToken
			}

			var dataLogin = customers.CustomerCore{
				ID:       data.ID.String(),
				Email:    data.Email,
				Password: data.Password,
			}

			return dataLogin, token, nil
		}
	}
	return customers.CustomerCore{}, "", errors.New("email pr password incorrect")
}

// Register implements entity.CustomerDataInterface.
func (customerRep *customerRepository) Register(data customers.CustomerCore) (row int, err error) {

	newUUID, UUIDerr := uuid.NewRandom()
	if UUIDerr != nil {
		return 0, UUIDerr
	}

	hashPassword, err := crypt.HashPassword(data.Password)
	if err != nil {
		return 0, err
	}

	var input = model.Customers{
		ID:       newUUID,
		Name:     data.Name,
		Email:    data.Email,
		Password: string(hashPassword),
		Address:  data.Address,
	}

	var checkEmail model.Customers
	dbCheck := customerRep.db.Where("email = ?", data.Email).First(&checkEmail)
	if dbCheck.Error != nil {
		if errors.Is(dbCheck.Error, gorm.ErrRecordNotFound) {
			erruser := customerRep.db.Save(&input)
			if erruser.Error != nil {
				return 0, erruser.Error
			}
			return 1, nil
		} else {
			return 0, dbCheck.Error
		}
	}

	return 0, errors.New("email already registered")
}
