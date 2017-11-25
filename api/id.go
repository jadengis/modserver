package api

import (
	"github.com/jadengis/modserver/duid"
)

var (
	idGenerator duid.IdGenerator
)

// SetIdGenerator sets the IdGenerator for the api package to use.
func SetIdGenerator(generator duid.IdGenerator) {
	idGenerator = generator
}

// Id is the base identification element to be used by all api elements.
type Id int

// IdEntity is a type containing a unique Id.
//
// Id returns the id of the IdEntity.
// SetId sets the id of the IdEntity.
type IdEntity interface {
	Id() Id
	SetId(Id)
}

// newId creates a new unique Id to be used by an IdEntity.
func newId() Id {
	return Id(idGenerator.NewId())
}
