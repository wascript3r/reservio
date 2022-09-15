package validator

type Rule[T any] interface {
	Name() string
	IsValid(d T) bool
}
