package post-data

import (
	"net/http"
	"github.com/TIBCOSoftware/flogo-lib/flow/activity"
	"github.com/op/go-logging"
)

// log is the default package logger
var log = logging.MustGetLogger("activity-tibco-post-data")

const (
	methodPost = "POST"
	
	ivURI     = "uri"
	ivPort    = "port"
	ivMethod  = "method"
	ivData    = "data"
	
	ovResult  = "result"
)

var validMethods = []string{methodPOST}

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
	
	
	method := strings.ToUpper(context.GetInput(ivMethod).(string))
	uri := context.GetInput(ivURI).(string)
	port := context.GetInput(ivPort).(string)
	urlData := url.Values{}
	urlData.Set("data", url:port)

	log.Debugf("HTTP Post: [%s] %s\n", method, uri)
	
	var reqBody io.Reader

	if method == methodPOST {

		content := context.GetInput(ivData)
		if context != nil {
			if str, ok := data.(string); ok {
				reqBody = bytes.NewBuffer([]byte(str))
			} else {
				b, _ := json.Marshal(data) //todo handle error
				reqBody = bytes.NewBuffer([]byte(b))
			}
		}
	} else {
		reqBody = nil
	}
	
	
	req, err := http.NewRequest(method, uri, reqBody)
	if reqBody != nil {
		req.Header.Set("Content-Type", "text/plain")
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	log.Debug("response Status:", resp.Status)
	respBody, _ := ioutil.ReadAll(resp.Body)

	var result interface{}

	d := json.NewDecoder(bytes.NewReader(respBody))
	d.UseNumber()
	err = d.Decode(&result)

	//json.Unmarshal(respBody, &result)

	if log.IsEnabledFor(logging.DEBUG) {
		log.Debug("response Body:", result)
	}

	context.SetOutput(ovResult, result)

	return true, nil
	
}