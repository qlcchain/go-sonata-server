/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package handler

import (
	"testing"

	"github.com/qlcchain/go-sonata-server/util"
)

func TestCancelProductOrderCreate(t *testing.T) {
	t.Log(util.ToString(CancelProductOrderCreate()))
}
