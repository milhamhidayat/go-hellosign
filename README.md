# go-hellosign

hellosign package to automate signature process when LA signature process. This is my attempt to create HelloSign wrapper in Go because there is no official package for Go.

## Configuration
a. Clone **go-hellosign** package

b. In your Go application, import this package and pass your api key to initialize it

```go
	
var helloSignAPI = hellosign.NewHelloSign("your_api_key")
```

(**NOTE** : if you want to test HelloSign callback in localhost)
c. Install ngrok in your machine

d. run ngrok with same port as your application so HelloSign can access your application in localhost

```sh
ngrok http yout_application_port
```

## Usage
### Send signature request

```go
payload := hellosign.SendSignatureRequestPayload{
		TestMode: 1,
		Title:    "LA/PO11/13",
		Subject:  "test subject",
		Message:  "test message",
		SignerEmail: []hellosign.SignerEmailDetail{
			hellosign.SignerEmailDetail{
				EmailAddress: "john@gmail.com",
				Name:         "john doe",
				Order:        0,
			},
			hellosign.SignerEmailDetail{
				EmailAddress: "rose@gmail.com",
				Name:         "rose smith",
				Order:        1,
			},
		},
		FileURL: []string{"[your_pdf_link_file]"},
	}
res, err := helloSignAPI.SendSignatureRequest(payload)
```

#### **NOTE**
Order of object is very important when to make request. This is the rule :
1. First signer email : order value must be 0
2. Second signer email : order value must be 1

### Get Signature Request Detail

```go
res, err := hs.GetSignatureRequestDetail("signature_request_id")
```

### Verify event callback hash
When you successfully send a signature request, HelloSign API will send event callback (event_time, event_type, event_hash). You need to verify event_hash to make sure the callback is valid. If valid, you need to send response with status 200 and a response body containing the following text: **Hello API Event Received**

```go
check := helloSignAPI.VerifyHash(eventTime, eventType, eventHash)
```

## Example

```go
package main

import (
	"encoding/json"
	"log"
	"net/http"

	hellosign "github.com/milhamhidayat/go-hellosign"

	"github.com/gorilla/mux"
)

var helloSignAPI = hellosign.NewHelloSign("api_key")

func sendSignature(w http.ResponseWriter, r *http.Request) {
	payload := hellosign.SendSignatureRequestPayload{
		TestMode: 1,
		Title:    "LA/PO11/13",
		Subject:  "test subject",
		Message:  "test message",
		SignerEmail: []hellosign.SignerEmailDetail{
			hellosign.SignerEmailDetail{
				EmailAddress: "john@gmail.com",
				Name:         "john doe",
				Order:        0,
			},
			hellosign.SignerEmailDetail{
				EmailAddress: "rose@gmail.com",
				Name:         "rose smith",
				Order:        1,
			},
		},
		FileURL: []string{"[your_pdf_link_file]"},
	}
	res, err := helloSignAPI.SendSignatureRequest(payload)
	w.Header().Set("Content-Type", "application/json")
	if err != nil {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(res)
	}
}

func handleCallback(w http.ResponseWriter, r *http.Request) {
	var eventResponse hellosign.CallbackEvent
	w.Header().Set("Content-Type", "application/json")

	r.ParseMultipartForm(0)

	event := r.FormValue("json")

	err := json.Unmarshal([]byte(event), &eventResponse)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err)
	}

	eventDetail := eventResponse.EventDetail

	check := helloSignAPI.VerifyHash(eventDetail.EventTime, eventDetail.EventType, eventDetail.EventHash)

	if check == true {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Hello API Event Received")
	} else {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode("Not Received")
	}
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/LA", sendSignature).Methods("POST")
	r.HandleFunc("/callback", handleCallback).Methods("POST")
	log.Fatal(http.ListenAndServe(":3000", r))
}

```


