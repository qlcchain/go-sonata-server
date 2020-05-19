/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package schema

import (
	"encoding/json"
	"reflect"
	"testing"

	"github.com/brianvoe/gofakeit/v5"
	"github.com/go-openapi/swag"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

func TestRelatedParty_MarshalJSON(t *testing.T) {
	m := &models.RelatedParty{
		AtReferredType:  "Organization",
		EmailAddress:    gofakeit.Email(),
		ID:              util.NewIDPtr(),
		Name:            swag.String(gofakeit.Name()),
		Number:          gofakeit.Phone(),
		NumberExtension: "",
		Role: []string{
			"Buyer", "Seller", "Site Contact",
		},
	}

	s := util.ToString(m)

	to := &RelatedParty{}
	data := []byte(s)
	if err := json.Unmarshal(data, to); err != nil {
		t.Fatal(err)
	} else {
		if to.Roles == "" {
			t.Fatal("invalid roles")
		}
		if !reflect.DeepEqual(m.Role, to.Role) {
			t.Fatalf("invalid role, exp: %v, got: %v", m.Role, to.Role)
		}
	}
}
