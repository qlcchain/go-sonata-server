package util

import "encoding/json"

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
