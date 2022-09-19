package models

type Role string

const (
	ClientRole  Role = "client"
	CompanyRole Role = "company"
	AdminRole   Role = "admin"
)

type User struct {
	ID       string
	Email    string
	Password string
	Role     Role
}
