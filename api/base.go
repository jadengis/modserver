package api

// BaseEntity is the base interface for persisted entities.
//
// IdEntity is embedded in every BaseEntity.
// TableName should return the database table name of the BaseEntity.
type BaseEntity interface {
	IdEntity
	TableName() string
}
