package api

// Name is a the type for names in entities.
type Name string

// NameEntity is the type for name holding entities.
//
// BaseEntity is embedded with the NameEntity.
// Name returns the name of the named entity.
// SetName sets the name of the named entity.
type NameEntity interface {
	BaseEntity
	Name() Name
	SetName(Name)
}
