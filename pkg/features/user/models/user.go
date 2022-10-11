package models

type Role string

const (
	ClientRole  Role = "client"
	CompanyRole Role = "company"
	AdminRole   Role = "admin"
)

func (r Role) String() string {
	return string(r)
}

type User struct {
	ID       string
	Email    string
	Password string
	Role     Role
}
