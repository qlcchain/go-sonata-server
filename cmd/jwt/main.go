package main

import (
	"encoding/base64"
	"os"

	"github.com/sirupsen/logrus"

	flags "github.com/jessevdk/go-flags"

	"github.com/qlcchain/go-sonata-server/auth"
)

var opts struct {
	Key   string   `short:"K" long:"key" description:"private key"`
	Roles []string `long:"roles" description:"permission roles" default:"admin"`
}

func main() {
	logrus.SetLevel(logrus.DebugLevel)
	if _, err := flags.ParseArgs(&opts, os.Args); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	logrus.Debugln(opts)

	claims := auth.NewRoleClaims(opts.Roles)
	key := auth.DefaultKey
	if opts.Key != "" {
		if data, err := base64.StdEncoding.DecodeString(opts.Key); err != nil {
			logrus.Fatal(err)
		} else {
			key = string(data)
		}
	}
	privateKey := auth.Decode(key)
	if token, err := auth.NewToken(claims, privateKey); err == nil {
		logrus.Infof("Authorization: Bearer %s", token)
	} else {
		logrus.Fatalf("failed to generate JWT token, %s", err)
	}
}
