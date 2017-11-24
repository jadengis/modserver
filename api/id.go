package api

// Id is the base identification element to be used by all api elements.
type Id int

// IdEntity is an interface for Id holding api elements.
type IdEntity interface {
	GetId() Id
}

// generateId creates a new unique Id to be used by an api element.
func generateId() Id {

}
