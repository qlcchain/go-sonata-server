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

type DebugCfg struct {
	Verbose bool `json:"verbose" short:"V" long:"verbose" description:"Show verbose debug information"`
	Mock    bool `json:"mock"  long:"mock" description:"Mock sample data"`
}

type ServerCfg struct {
	Domain         string   `json:"domain" long:"domain" description:"bind domain" default:"localhost"`
	AllowedOrigins []string `json:"allowedOrigins" long:"allowedOrigins" description:"AllowedOrigins of CORS" default:"*"`
}

type SonataCfg struct {
	Server *ServerCfg `json:"server"`
	Jwt    *JwtCfg    `json:"jwt"`
	Debug  *DebugCfg  `json:"debug"`
}

type JwtCfg struct {
	Key        string            `json:"key" short:"K" long:"key" description:"private key(elliptic P521) to sign JWT token"`
	PrivateKey *ecdsa.PrivateKey `json:"-"`
	PublicKey  *ecdsa.PublicKey  `json:"-"`
}

func (s *SonataCfg) String() string {
	return util.ToString(s)
}

var (
	Cfg = &SonataCfg{
		Server: &ServerCfg{},
		Jwt:    &JwtCfg{},
		Debug:  &DebugCfg{},
	}
)
