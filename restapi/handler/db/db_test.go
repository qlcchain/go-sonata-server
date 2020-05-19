/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package db

import (
	"fmt"
	"testing"

	"github.com/qlcchain/go-sonata-server/restapi/handler"

	"github.com/qlcchain/go-sonata-server/schema"

	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/util"
)

func setupTestCase(t *testing.T) (func(t *testing.T), *gorm.DB) {
	file := fmt.Sprintf("file:%s?mode=memory&cache=shared", util.NewID())
	var err error
	Store, err = gorm.Open("sqlite3", file)
	if err != nil {
		t.Fatal(err)
	}
	db := Store
	db.LogMode(false)
	if err = CreateTables(); err != nil {
		t.Fatal(err)
	}
	return func(t *testing.T) {
		if err := db.Close(); err != nil {
			t.Error(err)
		}
	}, db
}

func TestLogAddress(t *testing.T) {
	t.Skip()
	var address []*schema.GeographicAddress
	var site []*schema.GeographicSite
	for i := 0; i < 5; i++ {
		a, s := handler.GeographicSiteAndAddress()
		address = append(address, schema.FromGeographicAddress(a))
		site = append(site, schema.FromGeographicSite(s))
	}
	fmt.Println(util.ToString(address))
	fmt.Println(util.ToString(site))
}
