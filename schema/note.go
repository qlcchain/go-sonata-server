/*
 * Copyright (c) 2020. QLC Chain Team
 *
 * This software is released under the MIT License.
 * https://opensource.org/licenses/MIT
 */

package schema

import (
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/rs/xid"

	"github.com/qlcchain/go-sonata-server/models"
	"github.com/qlcchain/go-sonata-server/util"
)

type Note struct {
	ID *string `json:"id"`
	// The author of the note
	// Required: true
	Author *string `json:"author"`

	// The date of the note. Format is YYYY-MM-DDThh:mmTZD (e.g. 1997-07-16T19:20+01:00).
	// Required: true
	// Format: date-time
	Date *strfmt.DateTime `json:"date"`

	// The text of the note
	// Required: true
	Text *string `json:"text"`
}

func FromNotes(notes []*models.Note) []*Note {
	var r []*Note
	for _, n := range notes {
		to := &Note{}
		if err := util.Convert(n, to); err == nil {
			if to.ID == nil {
				to.ID = swag.String(xid.New().String())
			}
			r = append(r, to)
		}
	}
	return r
}

func ToNotes(notes []*Note) []*models.Note {
	var to []*models.Note
	_ = util.Convert(notes, &to)
	return to
}
