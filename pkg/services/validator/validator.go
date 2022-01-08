package validator

type Validator interface {
	Struct(s interface{}) error
}
