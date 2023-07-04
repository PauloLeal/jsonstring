package jsonstring

import (
	"encoding/json"
	"strconv"
	"testing"
)

type testStruct struct {
	Js T `json:"js"`
}

var testJsonString string = `
			{
				"js": {
					"fieldString1": "abc",
					"fieldNumber1": 123
				}
			}
		`

func TestUnmarshalJSON(t *testing.T) {
	var ts testStruct

	json.Unmarshal([]byte(testJsonString), &ts)

	d := ts.Js.Data()

	d1 := d["fieldString1"]
	if d1 != "abc" {
		t.Errorf("expected fieldString1 to be %s. Got %s.", "abc", d1)
	}

	d2 := d["fieldNumber1"]
	if d2 != float64(123) {
		t.Errorf("expected fieldNumber1 to be %d. Got %d.", 123, d2)
	}
}

func TestMarshalJSON(t *testing.T) {
	ts := testStruct{}
	ts.Js = `{"fieldString1":"abc","fieldNumber1":123}`

	s, err := json.Marshal(ts)
	if err != nil {
		t.Errorf(err.Error())
	}

	stjs, _ := strconv.Unquote(string(testJsonString))
	ss, _ := strconv.Unquote(string(s))
	if ss != stjs {
		t.Errorf("expected s to be %s. Got %s.", testJsonString, s)
	}

}
