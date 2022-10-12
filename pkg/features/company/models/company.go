package models

type Company struct {
	ID          string
	Name        string
	Address     string
	Description string
}

type CompanyInfo struct {
	Email string
	Company
	Approved bool
}

type CompanyUpdate struct {
	Name        *string
	Address     *string
	Description *string
	Approved    *bool
}

func (c *CompanyUpdate) IsEmpty() bool {
	return c.Name == nil && c.Address == nil && c.Description == nil && c.Approved == nil
}
