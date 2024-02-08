package entity

type CustomerDataInterface interface{
	Register(data CustomerCore) (row int, err error)
	Login(email, password string) (CustomerCore, string, error)
}

type CustomerUseCaseInterface interface{
	Register(data CustomerCore) (row int, err error)
	Login(email, password string) (CustomerCore, string, error)
}