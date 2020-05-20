/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package db

import (
	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/schema"
	"github.com/qlcchain/go-sonata-server/util"
)

func GetProductOfferingQualification(id string) (*models.ProductOfferingQualification, error) {
	poq := &schema.ProductOfferingQualification{}
	tx := GetTx()
	if err := tx.Where("id=?", id).First(poq).Error; err != nil {
		return nil, err
	}
	to := &models.ProductOfferingQualification{}
	if err := util.Convert(poq, to); err == nil {
		return to, nil
	} else {
		return nil, err
	}
}
