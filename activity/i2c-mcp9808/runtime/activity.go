package mcp9808

import (
	"strconv"
	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/op/go-logging"
	"github.com/d2r2/go-i2c"
)

// log is the default package logger
var log = logging.MustGetLogger("activity-tibco-rest")

const (
	Addr     = 0x18
	Bus      = 1 // Depends on Raspberry Pi version (0 or 1)
	RegTemp  = 0x05
	RegManu  = 0x06
	RegDevID = 0x07
	regResult   = "result"
)

type I2CActivity struct {
	metadata *activity.Metadata
}

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(&I2CActivity{metadata: md})
}

// Metadata returns the activity's metadata
func (a *I2CActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *I2CActivity) Eval(context activity.Context) (done bool, err error) {
	i, err := i2c.NewI2C(Addr, Bus)
	if err != nil {
		log.Fatal(err)
	}
	defer i.Close()

	manu, err := i.ReadRegU16BE(RegManu)
	if err != nil {
		log.Fatal(err)
	}
	devid, err := i.ReadRegU16BE(RegDevID)
	if err != nil {
		log.Fatal(err)
	}
	if manu != 0x54 || devid != 0x0400 {
		log.Fatal("MCP9808 device not found")
	}

	t, err := i.ReadRegU16BE(RegTemp)
	if err != nil {
		log.Fatal(err)
	}
        
	val := float32(t&0x0FFF)/float32(16)
	val2 := val * 9.0 / 5.0 + 32
	context.SetOutput(regResult, strconv.FormatFloat(float64(val2), 'f', 1, 32))
	return true, nil
	
}
