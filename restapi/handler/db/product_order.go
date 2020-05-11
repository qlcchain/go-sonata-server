/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package db

import (
	"github.com/jinzhu/gorm"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/schema"
	"github.com/qlcchain/go-sonata-server/util"
)

func GetProductOrder(db *gorm.DB, id string) (*models.ProductOrder, error) {
	po := &schema.ProductOrder{}
	if err := db.Set(AutoPreLoad, true).Where("id=?", id).First(po).Error; err != nil {
		return nil, err
	}
	to := &models.ProductOrder{}
	if err := util.Convert(po, to); err == nil {
		return to, nil
	} else {
		return nil, err
	}
}
