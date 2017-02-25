package postdata

import (
	"net/http"
         "net/url"
         "strings"
	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/op/go-logging"
)

// log is the default package logger
var log = logging.MustGetLogger("activity-tibco-post-data")

const (
	//methodPost = "POST"
	
	ivURI     = "uri"
	//ivPort    = "port"
	//ivMethod  = "method"
	ivData    = "data"
	
	//ovResult  = "result"
)

var validMethods = []string{methodPost}

type PostActivity struct {
	metadata *activity.Metadata
}

// init create & register activity
func init() {
	md := activity.NewMetadata(jsonMetadata)
	activity.Register(&PostActivity{metadata: md})
}

// Metadata returns the activity's metadata
func (a *PostActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *PostActivity) Eval(context activity.Context) (done bool, err error) {
	
	 
	 urlData := url.Values{}
         urlData.Set("temp", data)
         uri := uri + "?temp="

         resp, err := http.Post(uri, "text/plain", strings.NewReader(urlData.Encode()))

         //if err != nil {
         //        fmt.Println(err)
        // }

         //fmt.Println("Status : ", resp.Status)

	//context.SetOutput(ovResult, result)

	return true, nil
	
}