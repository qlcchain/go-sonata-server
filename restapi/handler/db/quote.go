/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package db

import (
	"github.com/qlcchain/go-sonata-server/schema"
)

func GetQuote(id string) (*schema.Quote, error) {
	q := &schema.Quote{}
	tx := GetTx()
	if err := tx.Where("id=?", id).First(q).Error; err != nil {
		return nil, err
	} else {
		return q, nil
	}
}
