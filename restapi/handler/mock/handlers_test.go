package mock

import (
	"testing"

	"github.com/qlcchain/go-sonata-server/restapi/handler/db"
	"github.com/qlcchain/go-sonata-server/util"
)

func Test_mockGeographicAddress(t *testing.T) {
	mockGeographicAddress()
	var addresses []db.GeographicAddressModel
	if err := DB.Preload("FieldedAddress").Find(&addresses).Error; err != nil {
		t.Fatal(err)
	}
	t.Log(len(addresses))
	for idx, a := range addresses {
		t.Log(idx, util.ToIndentString(a))
	}
}
