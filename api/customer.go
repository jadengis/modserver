package api

// Customer represents a customer in the application.
type Customer struct {
	id   Id     `json: "id"`
	Name string `json: "name"`
}

func (c Customer) GetId() Id {
	return id
}
