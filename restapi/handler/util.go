/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package handler

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/qlcchain/go-sonata-server/schema"

	log "github.com/sirupsen/logrus"

	"github.com/qlcchain/go-sonata-server/util"

	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"

	"github.com/qlcchain/go-sonata-server/models"
)

var (
	jsonContentType = "application/json;charset=utf-8"
	client          = &http.Client{}
)

func NewDatetime(dt time.Time) *strfmt.DateTime {
	t := strfmt.DateTime(dt)
	return &t
}

func NewDate(dt time.Time) *strfmt.Date {
	t := strfmt.Date(dt)
	return &t
}

func HrefToID(prefix, id string) string {
	return fmt.Sprintf("%s/%s", prefix, id)
}

func HttpPost(url string, v interface{}) (string, error) {
	content := util.ToString(v)
	log.Debug(content)
	resp, err := client.Post(url, jsonContentType, bytes.NewBuffer([]byte(content)))
	if err != nil {
		return "", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			log.Error(err)
		}
	}()

	if all, err := ioutil.ReadAll(resp.Body); err != nil {
		return "", err
	} else {
		return string(all), nil
	}
}

func VerifyField(v interface{}) (string, bool) {
	var value string
	switch t := v.(type) {
	case *string:
		value = swag.StringValue(t)
	case *strfmt.DateTime:
		if t != nil {
			value = t.String()
		}
	}

	return value, value != ""
}

func ParseType(v string) string {
	if strings.Contains(v, "=") {
		if split := strings.Split(v, "="); len(split) == 2 {
			return strings.Trim(split[1], " ")
		}
	}
	return v
}

func ToErrorRepresentation(principal *models.Principal) *models.ErrorRepresentation {
	if principal.Code != 0 {
		return &models.ErrorRepresentation{
			Code:   swag.Int32(principal.Code),
			Reason: swag.String(principal.Reason),
		}
	}
	return nil
}

func VerifyCallback(callback *string) *models.ErrorRepresentation {
	if _, err := url.Parse(*callback); err != nil {
		return &models.ErrorRepresentation{
			Code:   swag.Int32(24),
			Reason: swag.String("Invalid  body field"),
		}
	}
	return nil
}

func ConvertRelatedParty(from []*models.RelatedParty) []*schema.RelatedParty {
	if from == nil {
		return nil
	}
	var to []*schema.RelatedParty
	for _, i := range from {
		m := &schema.RelatedParty{}
		to = append(to, m.From(i))
	}
	return to
}
