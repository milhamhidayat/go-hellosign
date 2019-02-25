package hellosign

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

func prepareData(payload SendSignatureRequestPayload) string {
	testModeKey := "test_mode"
	titleKey := "title"
	subjectKey := "subject"
	messageKey := "message"
	useTextTagsKey := "use_text_tags"

	testMode := generateHelloSignData(testModeKey, payload.TestMode)
	title := generateHelloSignData(titleKey, payload.Title)
	subject := generateHelloSignData(subjectKey, payload.Subject)
	message := generateHelloSignData(messageKey, payload.Message)
	signer := generateSigners(payload.SignerEmail)
	fileURL := generateFileURL(payload.FileURL)
	useTextTags := generateHelloSignData(useTextTagsKey, "1")
	ccEmail := generateCCEmail(payload.CCEmail)

	signatureRequestPayload := fmt.Sprintf("{%v,%v,%v,%v,%v,%v,%v,%v}", testMode, title, subject, message, signer, fileURL, useTextTags, ccEmail)
	// signatureOptions := `"signing_options":"{"draw":true,"type":true,"upload":true,"phone":false,"default":type}"`

	finalSignatureRequestPayload := fmt.Sprintf(`%v`, signatureRequestPayload)

	return finalSignatureRequestPayload

}

func validateSignatureRequestPayload(payload SendSignatureRequestPayload) error {
	if len(payload.FileURL) == 0 {
		return errors.New("FileURL can't be empty")
	}

	for _, fileURL := range payload.FileURL {
		if fileURL == "" {
			return errors.New("FileURL can't be empty string")
		}
	}

	if payload.TestMode > 1 || payload.TestMode < 0 {
		return errors.New("test_mode value can only be 0 or 1")
	}

	if payload.Title == "" {
		return errors.New("title can't be empty")
	}

	if payload.Message == "" {
		return errors.New("message can't be empty")
	}

	if payload.Subject == "" {
		return errors.New("subject can't be empty")
	}

	if len(payload.SignerEmail) == 0 {
		return errors.New("signer detail in payload can't be empty")
	}

	if len(payload.CCEmail) == 0 {
		return errors.New("cc email can't be empty")
	}

	for _, ccEmail := range payload.CCEmail {
		if ccEmail == "" {
			return errors.New("cc email value can't be empty string")
		}
	}

	return nil
}

func generateCCEmail(CCEmail []string) string {
	var tmpCCEmail []string
	var ccEmailKey string

	for k, v := range CCEmail {
		ccEmailKey = fmt.Sprintf(`"cc_email_addresses[%d]":"%v"`, k, v)
		tmpCCEmail = append(tmpCCEmail, ccEmailKey)
	}

	finalCCEmail := strings.Join(tmpCCEmail, ",")
	return finalCCEmail
}

func generateSigners(signer []SignerEmailDetail) string {
	var tmpSignersList []string

	for key, value := range signer {
		s := structToMap(value)
		for index, data := range s {
			curlKey := fmt.Sprintf(`"signers[%d][%s]":"%v"`, key, index, data)
			tmpSignersList = append(tmpSignersList, curlKey)
		}
	}

	signersList := strings.Join(tmpSignersList, ",")
	return signersList
}

func generateFileURL(fileURL []string) string {
	var tmpFileURL []string
	var fileURLKey string

	for k, v := range fileURL {
		fileURLKey = fmt.Sprintf(`"file_url[%d]":"%v"`, k, v)
		tmpFileURL = append(tmpFileURL, fileURLKey)
	}

	finalFileURL := strings.Join(tmpFileURL, ",")
	return finalFileURL
}

func structToMap(signer SignerEmailDetail) map[string]interface{} {
	var signerMap map[string]interface{}
	tmpSignerMap, _ := json.Marshal(signer)
	json.Unmarshal(tmpSignerMap, &signerMap)
	return signerMap
}

func generateSignerKey(key string) string {
	var signerKey string
	switch key {
	case "EmailAddress":
		signerKey = "email_address"
	case "Name":
		signerKey = "name"
	case "Order":
		signerKey = "order"
	}
	return signerKey
}

func generateHelloSignData(key string, value interface{}) string {
	return fmt.Sprintf(`"%v":"%v"`, key, value)
}
