package mock

import (
	"testing"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/restapi/handler/db"

	"github.com/qlcchain/go-sonata-server/util"
)

func Test_mockGeographicAddress(t *testing.T) {
	mockGeographicAddress(2)
	if listGeographicAddress, err := db.ListGeographicAddress(DB); err != nil {
		t.Fatal(err)
	} else {
		if len(listGeographicAddress) > 0 {
			t.Log(len(listGeographicAddress))
			for idx, a := range listGeographicAddress {
				t.Log(idx, util.ToIndentString(a))
			}

			if address, err := db.GetGeographicAddress(DB, listGeographicAddress[0].ID); err != nil {
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
			if err := DB.Set(db.AutoPreLoad, true).Model(&fa).Related(&address, "ID").Error; err != nil {
				t.Fatal(err)
			} else {
				t.Log(util.ToString(address))
			}
		} else {
			t.Fatal("invalid address len,exp: 10, got: 0")
		}
	}
}
