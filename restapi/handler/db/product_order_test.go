/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package db

import (
	"testing"

	"github.com/qlcchain/go-sonata-server/restapi/handler"
	"github.com/qlcchain/go-sonata-server/util"
)

func TestGetProductOrderByParams(t *testing.T) {
	teardownTestCase, db := setupTestCase(t)
	defer teardownTestCase(t)

	create := handler.ProductOrderCreate()
	t.Log(util.ToString(create))
	if order, err := ProductOrderCreateToProductOrder(create); err == nil {
		if err := db.Create(order).Error; err != nil {
			t.Fatal(err)
		} else {
			t.Log(util.ToString(order))
		}
	} else {
		t.Fatal(err)
	}
}
