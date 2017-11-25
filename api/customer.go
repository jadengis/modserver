package api

const (
	// CustomerTableName is the name of the customer table.
	CustomerTableName = "customers"
)

// Customer represents a customer in the application.
type Customer struct {
	id   Id   `json: "id"`
	name Name `json: "name"`
}

// NewCustomer constructs a new Customer with all required properties.
func NewCustomer(name Name) *Customer {
	return &Customer{
		id:   newId(),
		name: name}
}

// TableName returns the name of the Customer table.
func (c *Customer) TableName() string {
	return CustomerTableName
}

// Id returns the id of the given Customer.
func (c *Customer) Id() Id {
	return c.id
}

// SetId sets the id of the given Customer.
func (c *Customer) SetId(id Id) {
	c.id = id
}

// Name returns the name of the given Customer.
func (c *Customer) Name() Name {
	return c.name
}

// SetName sets the name of the given Customer.
func (c *Customer) SetName(name Name) {
	c.name = name
}
