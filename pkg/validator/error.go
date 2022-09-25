package validator

type Error interface {
	error
	GetData() any
}
