package models

import (
	"github.com/jadengis/modserver/duid"
	"time"
)

var idGenerator duid.IdGenerator

func SetIdGenerator(generator duid.IdGenerator) {
	idGenerator = generator
}

type BaseModel struct {
	ID        uint       `sql:"type: bigint PRIMARY KEY" json:"id,string"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"deletedAt,omitempty"`
}
