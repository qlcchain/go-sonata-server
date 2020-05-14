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

	"github.com/go-openapi/swag"

	gs "github.com/qlcchain/go-sonata-server/restapi/operations/geographic_site"

	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/schema"
	"github.com/qlcchain/go-sonata-server/util"
)

func TestGetGeographicSiteByParams(t *testing.T) {
	teardownTestCase, db := setupTestCase(t)
	defer teardownTestCase(t)

	var sites []*schema.GeographicSite
	const size = 5
	for i := 0; i < size; i++ {
		a := handler.GeographicSite()
		to := schema.FromGeographicSite(a)
		sites = append(sites, to)
		//t.Log(util.ToString(to))
		db.Create(to)
	}

	if len(sites) == 0 {
		t.Fatal()
	}

	t.Log(util.ToString(sites))

	if site, err := GetGeographicSite(db, sites[0].ID); err == nil {
		if !reflect.DeepEqual(site, sites[0].To()) {
			t.Fatalf("invalid site, exp: %s, got: %s", util.ToString(sites[0].To()), util.ToString(site))
		}
	} else {
		t.Fatal(err)
	}

	from := sites[0]
	if r, err := GetGeographicSiteByParams(db, &gs.GeographicSiteFindParams{
		SiteCompanyName: swag.String(from.SiteCompanyName),
	}); err == nil {
		if !reflect.DeepEqual(r[0].To(), from.To()) {
			t.Fatalf("invalid site, exp: %s, got: %s", util.ToString(from.To()), util.ToString(r[0].To()))
		}
	} else {
		t.Fatal(err)
	}
}
