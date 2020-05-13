package schema

import (
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

type ProductTerm struct {
	ID *string `json:"id"`

	// Description of the commitment
	Description string `json:"description,omitempty"`

	// duration
	Duration *models.Quantity `json:"duration,omitempty" gorm:"embedded"`

	// Name of the commitment
	Name string `json:"name,omitempty"`

	// valid for
	ValidFor *models.TimePeriod `json:"validFor,omitempty" gorm:"embedded"`
}

func ToProductTerm(from []*ProductTerm) []*models.ProductTerm {
	var to []*models.ProductTerm
	if err := util.Convert(from, &to); err != nil {
		log.Error(err)
	}
	return to
}

func FromProductTerm(from []*models.ProductTerm) []*ProductTerm {
	if from == nil {
		return nil
	}
	var to []*ProductTerm
	if err := util.Convert(from, &to); err != nil {
		log.Error(err)
	}
	for _, term := range to {
		term.ID = util.NewOrDefaultPtr(term.ID)
	}
	return to
}
