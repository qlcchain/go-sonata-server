package schema

import (
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

type ProductRelationship struct {
	ID *string `json:"id"`

	// product
	// Required: true
	Product *ProductRef `json:"product" gorm:"foreignkey:ID"`

	// Indicates whether the type of relationship is "bundled", "reliesOn", or "comesFrom"
	// Required: true
	Type *string `json:"type"`
}

func ToProductRelationship(from []*ProductRelationship) []*models.ProductRelationship {
	var to []*models.ProductRelationship
	if err := util.Convert(from, &to); err != nil {
		log.Error(err)
	}
	return to
}

func FromProductRelationship(from []*models.ProductRelationship) []*ProductRelationship {
	var to []*ProductRelationship
	if err := util.Convert(from, &to); err != nil {
		log.Error(err)
	}
	for _, relationship := range to {
		relationship.ID = util.NewOrDefaultPtr(relationship.ID)
	}
	return to
}
