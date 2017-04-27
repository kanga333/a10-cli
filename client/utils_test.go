package client

import (
	"encoding/json"
	"testing"
)

const (
	falseData   = `{"boolean":0}`
	trueData    = `{"boolean":1}`
	illegalData = `{"boolean":100}`
)

func TestnumBool(t *testing.T) {
	type asBoolean struct {
		Boolean NumBool `json:"boolean"`
	}

	var asTrue asBoolean
	err := json.Unmarshal([]byte(trueData), &asTrue)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if asTrue.Boolean != true {
		t.Error("asTrue.Boolean should be true but", asTrue.Boolean)
	}
	trueByte, err := json.Marshal(&asTrue)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if string(trueByte) != trueData {
		t.Error("asTrue.Boolean should be true but", string(trueByte))
	}

	var asFalse asBoolean
	err = json.Unmarshal([]byte(falseData), &asFalse)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if asFalse.Boolean != false {
		t.Error("asFalse.Boolean should be true but", asFalse.Boolean)
	}
	falseByte, err := json.Marshal(&asFalse)
	if err != nil {
		t.Errorf("should not raise error: %v", err)
	}
	if string(falseByte) != falseData {
		t.Error("asTrue.Boolean should be true but", string(falseByte))
	}

	var asNil asBoolean
	err = json.Unmarshal([]byte(illegalData), &asNil)
	if err == nil {
		t.Error("should raise error")
	}

}
