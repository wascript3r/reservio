package models

type Company struct {
	UserID      string
	Name        string
	Address     string
	Description string
}

type CompanyInfo struct {
	Company
	Email string
}

type CompanyUpdate struct {
	Name        string
	Address     string
	Description string
}
