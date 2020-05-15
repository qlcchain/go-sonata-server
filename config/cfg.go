/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package config

import (
	"crypto/ecdsa"

	"github.com/qlcchain/go-sonata-server/util"
)

// TODO: add mock data option
type ServerCfg struct {
	Key        string            `json:"key" short:"K" long:"key" description:"private key(elliptic P521) to sign JWT token"`
	Verbose    bool              `json:"verbose" short:"V" long:"verbose" description:"Show verbose debug information"`
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	PublicKey  *ecdsa.PublicKey  `json:"-"`
}

func (s *ServerCfg) String() string {
	return util.ToString(s)
}

var (
	Cfg = &ServerCfg{}
)
