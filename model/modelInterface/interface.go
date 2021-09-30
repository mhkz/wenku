package modelInterface

type CURD interface {
	Create() (error, CURD)
	Updata() (error, CURD)
	Read() (error, CURD)
	Delete() (error, CURD)
}
