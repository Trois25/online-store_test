package service

import (
	customers "store/features/customers/entity"

	"errors"
	"regexp"
)

type customerUseCase struct {
	customerRepository customers.CustomerDataInterface
}

// Login implements entity.CustomerUseCaseInterfacec.
func (cu *customerUseCase) Login(email string, password string) (customers.CustomerCore, string, error) {
	if email == "" || password == "" {
		return customers.CustomerCore{}, "", errors.New("error, email or password can't be empty")
	}

	logindata, token, err := cu.customerRepository.Login(email, password)

	if err != nil {
		return customers.CustomerCore{}, "", err
	}


	return logindata, token, nil
}

// Register implements entity.CustomerUseCaseInterfacec.
func (cu *customerUseCase) Register(data customers.CustomerCore) (row int, err error) {
	if data.Email == "" || data.Password == "" {
		return 0, errors.New("error, email or password can't be empty")
	}

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailRegex, data.Email)
	if !match {
		return 0, errors.New("error. email format not valid")
	}

	errcustomerdata, err := cu.customerRepository.Register(data)
	if err != nil {
		return 0, err
	}

	return errcustomerdata, nil
}

func New(CustomerUseCase customers.CustomerDataInterface) customers.CustomerUseCaseInterface {
	return &customerUseCase{
		customerRepository: CustomerUseCase,
	}
}
