/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package mock

import (
	"testing"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler/db"

	"github.com/qlcchain/go-sonata-server/util"
)

func Test_mockGeographicAddress(t *testing.T) {
	mockGeographicAddress(2)
	if listGeographicAddress, err := db.ListGeographicAddress(Store); err != nil {
		t.Fatal(err)
	} else {
		if len(listGeographicAddress) > 0 {
			t.Log(len(listGeographicAddress))
			for idx, a := range listGeographicAddress {
				t.Log(idx, util.ToIndentString(a))
			}

			if address, err := db.GetGeographicAddress(Store, listGeographicAddress[0].ID); err != nil {
				t.Fatal(err)
			} else {
				if address == nil {
					t.Fatalf("invalid address by id %s", listGeographicAddress[0].ID)
				} else {
					t.Log(util.ToString(address))
				}
			}

			fa := listGeographicAddress[0].FieldedAddress
			var address models.GeographicAddress
			if err := Store.Set(db.AutoPreLoad, true).Model(&fa).Related(&address, "ID").Error; err != nil {
				t.Fatal(err)
			} else {
				t.Log(util.ToString(address))
			}
		} else {
			t.Fatal("invalid address len,exp: 10, got: 0")
		}
	}
}

func TestMockQuoteCreate(t *testing.T) {
	//models.ProductOfferingQualificationCreate{}
	val := ProductOfferingQualificationCreate()
	t.Log(util.ToString(val))

	//t.Log(util.ToString(&product_offering_qualification.ProductOfferingQualificationFindParams{}))
}

//
//type Account struct {
//	ID    string
//	Extra *Extra `json:"extra" gorm:"embedded"`
//}
//
//type Extra struct {
//	F1 string
//	F2 string
//}
//
//func TestEmbedded(t *testing.T) {
//	Store.AutoMigrate(&Account{})
//	id := xid.New().String()
//	Store.Create(&Account{
//		ID: id,
//		//Extra: &Extra{
//		//	F1: "f1",
//		//	F2: "f2",
//		//},
//	})
//	r := &Account{}
//	if err := Store.Set(db.AutoPreLoad, true).Where("id=?", id).First(r).Error; err != nil {
//		t.Fatal(err)
//	} else {
//		t.Log(util.ToIndentString(r))
//	}
//}
