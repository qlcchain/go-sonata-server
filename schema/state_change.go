package schema

import (
	"github.com/go-openapi/strfmt"

	"github.com/qlcchain/go-sonata-server/models"
)

type StateChange struct {
	ID *string `json:"id"`
	// change date
	// Format: date-time
	ChangeDate strfmt.DateTime `json:"changeDate,omitempty"`

	// change reason
	ChangeReason string `json:"changeReason,omitempty"`

	// state
	State models.ProductOfferingQualificationStateType `json:"state,omitempty"`
}
