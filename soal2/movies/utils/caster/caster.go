package caster

import "encoding/json"

func Cast(from interface{}, to interface{}) error {
	bt, e := json.Marshal(from)
	if e != nil {
		return e
	}

	e = json.Unmarshal(bt, to)
	if e != nil {
		return e
	}

	return nil
}
