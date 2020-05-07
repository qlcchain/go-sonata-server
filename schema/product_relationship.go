package schema

type ProductRelationship struct {
	ID *string `json:"id"`

	// product
	// Required: true
	Product *ProductRef `json:"product" gorm:"foreignkey:ID"`

	// Indicates whether the type of relationship is "bundled", "reliesOn", or "comesFrom"
	// Required: true
	Type *string `json:"type"`
}
