package validator

type Rule interface {
	Alias() string
	Tags() string
}
