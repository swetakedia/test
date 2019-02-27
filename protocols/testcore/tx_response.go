package testcore

const (
	// TXStatusError represents the status value returned by test-core when an error occurred from
	// submitting a transaction
	TXStatusError = "ERROR"

	// TXStatusPending represents the status value returned by test-core when a transaction has
	// been accepted for processing
	TXStatusPending = "PENDING"

	// TXStatusDuplicate represents the status value returned by test-core when a submitted
	// transaction is a duplicate
	TXStatusDuplicate = "DUPLICATE"
)

// TXResponse represents the response returned from a submission request sent to test-core's /tx
// endpoint
type TXResponse struct {
	Exception string `json:"exception"`
	Error     string `json:"error"`
	Status    string `json:"status"`
}

// IsException returns true if the response represents an exception response from test-core
func (resp *TXResponse) IsException() bool {
	return resp.Exception != ""
}
