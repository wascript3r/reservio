package models

type Client struct {
	ID        string
	FirstName string
	LastName  string
	Phone     string
}

type ClientInfo struct {
	Client
	Email string
}
