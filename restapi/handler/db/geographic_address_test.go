/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package db

import (
	"reflect"
	"testing"

	_ "github.com/jinzhu/gorm/dialects/sqlite"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/schema"
	"github.com/qlcchain/go-sonata-server/util"
)

func TestGetGeographicAddressByFieldedAddress(t *testing.T) {
	teardownTestCase, db := setupTestCase(t)
	defer teardownTestCase(t)

	var addresses []*schema.GeographicAddress
	const size = 1
	for i := 0; i < size; i++ {
		a := handler.GeographicAddress()
		to := &schema.GeographicAddress{
			ID: util.NewID(),
		}
		if err := util.Convert(a, to); err == nil {
			addresses = append(addresses, to)
			t.Log(util.ToString(to))
			db.Create(to)
		} else {
			t.Fatal(err)
		}
	}
	if len(addresses) == 0 {
		t.Fatal()
	}

	if address, err := ListGeographicAddress(db); err != nil {
		t.Fatal(err)
	} else if len(address) != len(addresses) {
		t.Fatalf("invalid address, exp: %d, got: %d", len(addresses), len(address))
	} else {
		t.Log(util.ToString(address))
	}

	fa := addresses[0].FieldedAddress
	if address, err := GetGeographicAddressByFieldedAddress(db, &models.FieldedAddressRequest{
		City:                 fa.City,
		Country:              fa.Country,
		GeographicSubAddress: nil,
		Locality:             fa.Locality,
		StreetName:           fa.StreetName,
	}); err != nil {
		t.Fatal(err)
	} else if len(address) != 1 {
		t.Fatalf("invalid address, exp: 1, got: %d", len(address))
	} else {
		if !reflect.DeepEqual(address[0], addresses[0]) {
			t.Fatalf("invalid schema address, exp: %v, got: %v", util.ToString(addresses[0]), util.ToString(address[0]))
		}
	}

	fa2 := addresses[0].FormattedAddress
	if address, err := GetGeographicAddressByFormattedAddress(db, &models.FormattedAddressRequest{
		AddrLine1: fa2.AddrLine1,
		City:      fa2.City,
	}); err != nil {
		t.Fatal(err)
	} else if len(address) != 1 {
		t.Fatalf("invalid address, exp: 1, got: %d", len(address))
	} else {
		if !reflect.DeepEqual(address[0], addresses[0]) {
			t.Fatalf("invalid schema address, exp: %v, got: %v", util.ToString(addresses[0]), util.ToString(address[0]))
		}
	}
}
