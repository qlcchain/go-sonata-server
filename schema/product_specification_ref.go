package schema

import (
	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

type ProductSpecificationRef struct {

	// describing
	Describing *models.Describing `json:"describing,omitempty" gorm:"embedded"`

	// A unique identifier of the product spec.
	// Required: true
	ID *string `json:"id"`
}

func (m *ProductSpecificationRef) To() *models.ProductSpecificationRef {
	to := &models.ProductSpecificationRef{}
	if err := util.Convert(m, to); err != nil {
		log.Error(err)
	}
	return to
}

func FromProductSpecificationRef(f *models.ProductSpecificationRef) *ProductSpecificationRef {
	if f == nil {
		return nil
	}
	p := &ProductSpecificationRef{}
	if err := util.Convert(f, p); err != nil {
		log.Error(err)
	}
	p.ID = util.NewOrDefaultPtr(p.ID)
	return p
}
