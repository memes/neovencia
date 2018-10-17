package neovencia_claims

import (
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
)

func (model *Claim) BeforeCreate(scope *gorm.Scope) error {
	uuid := uuid.New()
	return scope.SetColumn("Id", uuid.String())
}
