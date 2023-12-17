package main

import "fmt"

type Address struct {
	Street string
	Number string
}

type Person interface {
	Name() string
}

type Client struct {
	ID        int
	FirstName string
	LastName  string
	Active    bool
	Address
}

func (c Client) Name() string {
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

func FullName(person Person) string {
	return person.Name()
}

type Company struct {
	OfficialName string
}

func (c Company) Name() string {
	return c.OfficialName
}

func main() {
	client := Client{
		ID:        1986532,
		FirstName: "Jose",
		LastName:  "Silva",
		Active:    true,
	}

	client.Active = false
	client.Street = "Avenida Paulista"
	client.Address.Number = "1000A"

	FullName(client)

	fmt.Println(client)

	company := Company{}
	FullName(company)
}
