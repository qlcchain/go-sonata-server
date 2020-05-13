/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	flags "github.com/jessevdk/go-flags"

	"github.com/qlcchain/go-sonata-server/auth"
)

var opts struct {
	Key   string   `short:"K" long:"key" description:"private key"`
	Roles []string `long:"roles" description:"permission roles" default:"admin"`
}

func main() {
	log.SetLevel(log.DebugLevel)
	log.SetFormatter(&log.TextFormatter{})
	if _, err := flags.ParseArgs(&opts, os.Args); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	log.Debugln(opts)

	claims := auth.NewRoleClaims(opts.Roles)
	privateKey := auth.Decode(auth.DefaultKey)
	if opts.Key != "" {
		var err error
		if privateKey, err = auth.FromBase64(opts.Key); err != nil {
			log.Fatal(err)
		}
	}

	if token, err := auth.NewToken(claims, privateKey); err == nil {
		log.Infof("Authorization: Bearer %s", token)
	} else {
		log.Fatalf("failed to generate JWT token, %s", err)
	}
}
