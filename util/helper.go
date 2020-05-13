/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package util

import (
	"encoding/json"

	"github.com/go-openapi/swag"
	"github.com/rs/xid"
)

func ToString(v interface{}) string {
	bytes, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(bytes)
}

func ToIndentString(v interface{}) string {
	b, err := json.MarshalIndent(&v, "", "\t")
	if err != nil {
		return ""
	}
	return string(b)
}

func Convert(from, to interface{}) error {
	if data, err := json.Marshal(from); err != nil {
		return err
	} else {
		if err := json.Unmarshal(data, to); err != nil {
			return err
		} else {
			return nil
		}
	}
}

func NewID() string {
	return xid.New().String()
}

func NewIDPtr() *string {
	return swag.String(NewID())
}

func NewOrDefaultPtr(id *string) *string {
	if id == nil {
		return NewIDPtr()
	}
	return id
}

func NewOrDefault(id string) string {
	if id == "" {
		return NewID()
	}
	return id
}
