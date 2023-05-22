package domain

type Contact struct {
	ID          int
	FirstName   string
	LastName    string
	PhoneNumber string
}

func (c *Contact) FullName() string {
	return c.FirstName + " " + c.LastName
}
