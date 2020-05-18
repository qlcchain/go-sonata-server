/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package handler

import "github.com/qlcchain/go-sonata-server/restapi/operations"

type Binder interface {
	Bind(api *operations.SonataAPI)
}
