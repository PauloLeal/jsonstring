package jsonstring

import "encoding/json"

type T string

type jsonMap map[string]any

func (j *T) Data() jsonMap {
	var nm jsonMap
	b, _ := j.MarshalJSON()
	json.Unmarshal(b, &nm)
	return nm
}

func (j *T) MarshalJSON() ([]byte, error) {
	var n jsonMap
	json.Unmarshal([]byte(*j), &n)
	return json.Marshal(n)
}

func (j *T) UnmarshalJSON(b []byte) error {
	*j = T(b)
	return nil
}
