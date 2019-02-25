package hellosign

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"time"

	jsoniter "github.com/json-iterator/go"
	"github.com/parnurzeal/gorequest"
)

// APIHelloSign - List of hello sign api method
type APIHelloSign interface {
	SendSignatureRequest(payload SendSignatureRequestPayload) (SignatureRequestResponse, error)
	GetSignatureRequestDetail(signatureID string) (SignatureRequestResponse, error)
	CancelSignatureRequest(signatureID string) (string, error)
	VerifyHash(eventTime string, eventType string, eventHash string) bool
}

// helloSign - contain api key for authentication & authorization
type helloSign struct {
	apiKey string
}

// NewHelloSign - return helloSign struct with APIHelloSign interface
func NewHelloSign(apiKey string) APIHelloSign {
	return &helloSign{
		apiKey: apiKey,
	}
}

var helloSignURL = "https://api.hellosign.com/v3/signature_request/"

func (hs *helloSign) SendSignatureRequest(payload SendSignatureRequestPayload) (SignatureRequestResponse, error) {
	var successResponse SignatureRequestResponse
	var errorResponse ErrorResponse

	err := validateSignatureRequestPayload(payload)
	if err != nil {
		return successResponse, err
	}

	signatureRequestPayload := prepareData(payload)

	request := gorequest.New().Timeout(10*time.Second).SetBasicAuth(hs.apiKey, "")
	resp, body, errs := request.Post(helloSignURL + "/send").
		Type("multipart").
		Send(signatureRequestPayload).
		End()

	if errs != nil {
		return successResponse, errors.New("Failed to get response from HelloSign")
	}

	if resp.StatusCode != 200 {
		err := jsoniter.Unmarshal([]byte(body), &errorResponse)
		if err != nil {
			return successResponse, errors.New("Failed to unmarshal HelloSign error reponse")
		}
		return successResponse, errors.New(errorResponse.Error.ErrorMessage)

	}

	err = jsoniter.Unmarshal([]byte(body), &successResponse)
	if err != nil {
		return successResponse, errors.New("Failed to unmarshal HelloSign success reponse")
	}

	return successResponse, nil

}

func (hs *helloSign) GetSignatureRequestDetail(signatureID string) (SignatureRequestResponse, error) {
	var successResponse SignatureRequestResponse
	var errorResponse ErrorResponse

	if signatureID == "" {
		return successResponse, errors.New("Signature Request Id can't be empty")
	}

	request := gorequest.New().Timeout(10*time.Second).SetBasicAuth(hs.apiKey, "")
	resp, body, errs := request.Get(helloSignURL+signatureID).
		Type("json").
		Set("Content-Type", "application/json").
		End()

	if errs != nil {
		return successResponse, errors.New("Failed to get response from HelloSign")
	}

	if resp.StatusCode != 200 {
		err := jsoniter.Unmarshal([]byte(body), &errorResponse)
		if err != nil {
			return successResponse, errors.New("Failed to unmarshal HelloSign error reponse")
		}
		return successResponse, errors.New(errorResponse.Error.ErrorMessage)
	}

	err := jsoniter.Unmarshal([]byte(body), &successResponse)
	if err != nil {
		return successResponse, errors.New("Failed to unmarshal HelloSign success reponse")
	}

	return successResponse, nil
}

func (hs *helloSign) VerifyHash(eventTime string, eventType string, eventHash string) bool {
	data := eventTime + eventType

	dataByte := []byte(data)
	eventHashByte := []byte(eventHash)
	apiKeyByte := []byte(hs.apiKey)

	h := hmac.New(sha256.New, apiKeyByte)
	h.Write(dataByte)
	sha := hex.EncodeToString(h.Sum(nil))
	finalSha := []byte(sha)
	return hmac.Equal(finalSha, eventHashByte)
}

func (hs *helloSign) CancelSignatureRequest(signatureID string) (string, error) {
	var errorResponse ErrorResponse

	if signatureID == "" {
		return "", errors.New("Signature Request Id can't be empty")
	}

	request := gorequest.New().Timeout(10*time.Second).SetBasicAuth(hs.apiKey, "")
	resp, body, err := request.Post(helloSignURL+"cancel/"+signatureID).
		Type("json").
		Set("Content-Type", "application/json").
		End()

	if err != nil {
		return "", errors.New("Failed to get response from HeloSign")
	}

	if resp.StatusCode != 200 {
		err := jsoniter.Unmarshal([]byte(body), &errorResponse)
		if err != nil {
			return "", errors.New("Failed to unmarshal HelloSign error reponse")
		}
		return "", errors.New(errorResponse.Error.ErrorMessage)
	}

	return "Success cancel HelloSign signature request", nil
}
