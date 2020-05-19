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

func FindSubscriber(id string) (*schema.HubSubscriber, error) {
	var r = &schema.HubSubscriber{}
	err := Store.Where("id=?", id).First(r).Error
	return r, err
}

func ListSubscribers(atType string) ([]*schema.HubSubscriber, error) {
	var r []*schema.HubSubscriber
	err := Store.Where("type=?", atType).Find(&r).Error
	return r, err
}

func AddSubscriber(subscriber *schema.HubSubscriber) error {
	return Store.Create(subscriber).Error
}

func DeleteSubscriber(id string) error {
	return Store.Where("id=?", id).Delete(&schema.HubSubscriber{}).Error
}
