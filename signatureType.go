package hellosign

type (

	// SendSignatureRequestPayload - payload for send signature request
	SendSignatureRequestPayload struct {
		TestMode    int                 `json:"test_mode"`
		Title       string              `json:"title"`
		Subject     string              `json:"subject"`
		Message     string              `json:"message"`
		SignerEmail []SignerEmailDetail `json:"signers"`
		FileURL     []string            `json:"file_url"`
		CCEmail     []string            `json:"cc_email"`
	}

	// SignerEmailDetail - detail for signer
	SignerEmailDetail struct {
		EmailAddress string `json:"email_address"`
		Name         string `json:"name"`
		Order        int    `json:"order"`
	}

	// SignatureRequestResponse - signature request response from hello sign
	SignatureRequestResponse struct {
		SignatureRequest SignatureRequestDetail `json:"signature_request"`
	}

	// SignatureRequestDetail - signature request detail
	SignatureRequestDetail struct {
		SignatureRequestID    string               `json:"signature_request_id"`
		TestMode              bool                 `json:"test_mode"`
		Title                 string               `json:"title"`
		OriginalTitle         string               `json:"original_title"`
		Subject               string               `json:"subject"`
		Message               string               `json:"message"`
		IsComplete            bool                 `json:"is_complete"`
		IsDeclined            bool                 `json:"is_declined"`
		HasError              bool                 `json:"has_error"`
		ResponseData          []ResponseDataDetail `json:"response_data"`
		FilesURL              string               `json:"files_url"`
		RequesterEmailAddress string               `json:"requester_email_address"`
		Signers               []SignerDetail       `json:"signatures"`
	}

	// ResponseDataDetail - response data detail from hello sign callback
	ResponseDataDetail struct {
		Name        string `json:"name"`
		Type        string `json:"type"`
		Required    bool   `json:"required"`
		APIID       string `json:"api_id"`
		Value       string `json:"value"`
		SignatureID string `json:"signature_id"`
	}

	// SignerDetail - signer detail
	SignerDetail struct {
		SignatureID        string `json:"signature_id"`
		HasPin             bool   `json:"has_pin"`
		SignerEmailAddress string `json:"signer_email_address"`
		SignerName         string `json:"signer_name"`
		Order              int    `json:"order"`
		StatusCode         string `json:"status_code"`
		SignedAt           int64  `json:"signed_at"`
		LastViewedAt       int64  `json:"last_viewed_at"`
		LastRemindedAt     int64  `json:"last_reminded_at"`
		Error              string `json:"error"`
	}

	// ErrorResponse - error response from hello sign api
	ErrorResponse struct {
		Error ErrorDetail `json:"error"`
	}

	// ErrorDetail - error response detail
	ErrorDetail struct {
		ErrorMessage string `json:"error_msg"`
		ErrorName    string `json:"error_name"`
	}

	// CallbackEvent - callback event payload from hello sign
	CallbackEvent struct {
		EventDetail      EventDetail            `json:"event"`
		SignatureRequest SignatureRequestDetail `json:"signature_request"`
	}

	// EventDetail - event detail from hello sign callback
	EventDetail struct {
		EventType     string         `json:"event_type"`
		EventTime     string         `json:"event_time"`
		EventHash     string         `json:"event_hash"`
		EventMetadata MetaDataDetail `json:"event_metadata"`
	}

	// MetaDataDetail - event meta data detail
	MetaDataDetail struct {
		RelatedSignatureID   string `json:"related_signature_id"`
		ReportedForAccountID string `json:"reported_for_account_id"`
		ReportedForAppID     string `json:"reported_for_app_id"`
	}
)
