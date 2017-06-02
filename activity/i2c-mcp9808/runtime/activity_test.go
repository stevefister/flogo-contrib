package i2c-mcp9808

import (
	"testing"

	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/TIBCOSoftware/flogo-lib/flow/test"
)

func TestRegistered(t *testing.T) {
	act := activity.Get("i2c-mcp9808")

	if act == nil {
		t.Error("Activity Not Registered")
		t.Fail()
		return
	}
}

func TestReadState(t *testing.T) {

	act := activity.Get("i2c-mcp9808")

	tc := test.NewTestActivityContext(act.Metadata())

	//setup attrs
	tc.SetInput("method", "Read State")
	tc.SetInput("pin number", 3)
	//eval
	_, err := act.Eval(tc)
	if err != nil {
		log.Errorf("Error occured: %+v", err)
	}
	val := tc.GetOutput("result")
	log.Debugf("Resut %s", val)

}
