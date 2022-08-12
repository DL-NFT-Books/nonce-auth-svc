package apierrors

// feel free to use these codes, add your own or pass in your custom ones to methods directly as strings
const (
	// Handlers
	CodeEmptyBody      = "empty_body"
	CodeBadRequestData = "bad_request_data"
	// Login
	CodeNotRegistered = "not_registered"
	// Email Verification
	CodeInvalidEmailVerification = "invalid_email_verification"
	// Quotes
	CodeFailedPriceLookup = "failed_price_lookup"
	// Address Deletion
	CodeActiveAddressDeactivation = "active_address_deactivation"
	// Yoti
	CodeFailedYotiExchange = "failed_yoti_exchange"
	CodeFailedYotiVerification = "failed_yoti_verification"
	// Signature Verification
	CodeNonceNotFound               = "nonce_not_found"
	CodeSignatureVerificationFailed = "signature_verification_failed"
	// Session Authentication
	CodeSessionTokenNotFound = "session_token_not_found"
	CodeSessionTokenInvalid  = "session_token_invalid"
	// Resource Dependency
	CodeResourceNotFound = "resource_not_found"
	// Conflicts
	CodeAddressExists = "address_exists"
	// Suspended/Rejected User
	CodeUserSuspended = "user_suspended"
	CodeUserRejected = "user_rejected"
	// Verification Deletion
	CodeVerificationDeletionNotSupported = "verification_deletion_not_supported"
	// Faucets
	CodeFaucetAlreadyUsed = "faucet_already_used"
	// Not owned resources
	CodeResourceNotOwned = "resource_not_owned"
	CodeNotAddressOwner = "not_address_owner"
	// General Errors
	CodeBadRequest    = "bad_request"
	CodeUnauthorized  = "unauthorized"
	CodeForbidden     = "forbidden"
	CodeNotFound      = "not_found"
	CodeConflict      = "conflict"
	CodeTooEarly      = "too_early"
	CodeInternalError = "internal_error"
)
