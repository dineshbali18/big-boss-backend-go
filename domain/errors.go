package domain

import (
	"errors"
	"net/http"
)

// ResponseError struct holds error info to send in response
type ResponseError struct {
	ErrorCode        string `json:"errorCode"`
	ErrorDescription string `json:"errorDescription"`
	Status           int    `json:"-"`
}

func (re ResponseError) Error() string {
	return re.ErrorCode
}

// InvalidReqBody represents error code for invalid request body
const InvalidReqBody = "invalidRequestBody"

// UnexpectedError represents error code for unexpected error
const UnexpectedError = "unexpectedError"

// InvalidVariantValue represents error code for invalid variant value
const InvalidVariantValue = "invalidVariantValue"

// InvalidABTestID represents error code for invalid AB test id
const InvalidABTestID = "InvalidABTestID"

// DBError represents error code for database error
const DBError = "databaseError"

var (
	// InvalidAppVersion represents error message for invalidAppVersion
	InvalidAppVersion = ResponseError{"invalidAppVersion", "invalid appVersion provided", http.StatusBadRequest}
	// MissingAppVersion represents error message for missingAppVersion
	MissingAppVersion = ResponseError{"missingAppVersion", "appVersion missing in the input params", http.StatusBadRequest}

	// InvalidDeviceType represents error message for invalidDeviceType
	InvalidDeviceType = ResponseError{"invalidDeviceType", "invalid deviceType provided", http.StatusBadRequest}
	// MissingDeviceType represents error message for missingDeviceType
	MissingDeviceType = ResponseError{"missingDeviceType", "deviceType missing in the input params", http.StatusBadRequest}

	// InvalidDeviceId represents error message for invalidDeviceId
	InvalidDeviceId = ResponseError{"invalidDeviceId", "invalid deviceId provided", http.StatusBadRequest}
	// MissingDeviceId represents error message for missingDeviceId
	MissingDeviceId = ResponseError{"missingDeviceId", "deviceId missing in the input params", http.StatusBadRequest}

	// InvalidProductId represents error message for invalidProductId
	InvalidProductId = ResponseError{"invalidProductId", "invalid productId provided", http.StatusBadRequest}
	// MissingProductId represents error message for missingProductId
	MissingProductId = ResponseError{"missingProductId", "productId missing in the path params", http.StatusBadRequest}

	// InvalidId represents error message for invalidId
	InvalidId = ResponseError{"invalidId", "invalid Id provided", http.StatusBadRequest}
	// MissingId represents error message for missingId
	MissingId = ResponseError{"missingId", "Id missing in the path params", http.StatusBadRequest}

	// InvalidClientId represents error message for invalidClientId
	InvalidClientId = ResponseError{"invalidClientId", "invalid clientId provided", http.StatusBadRequest}
	// MissingClientId represents error message for missingClientId
	MissingClientId = ResponseError{"missingClientId", "clientId missing in the input params", http.StatusBadRequest}

	// InvalidPage represents error message for invalidPage
	InvalidPage = ResponseError{"invalidPage", "invalid page provided", http.StatusBadRequest}
	// MissingPage represents error message for missingPage
	MissingPage = ResponseError{"missingPage", "page missing in the input params", http.StatusBadRequest}

	// InvalidLimit represents error message for invalidLimit
	InvalidLimit = ResponseError{"invalidLimit", "invalid limit provided", http.StatusBadRequest}
	// MissingLimit represents error message for missingLimit
	MissingLimit = ResponseError{"missingLimit", "limit missing in the input params", http.StatusBadRequest}

	// InvalidCountryCode represents error message for invalidCountryCode
	InvalidCountryCode = ResponseError{"invalidCountryCode", "invalid countryCode provided", http.StatusBadRequest}
	// MissingCountryCode represents error message for missingCountryCode
	MissingCountryCode = ResponseError{"missingCountryCode", "countryCode missing in the input params", http.StatusBadRequest}

	// InvalidRegionCode represents error message for invalidRegionCode
	InvalidRegionCode = ResponseError{"invalidRegionCode", "invalid regionCode provided", http.StatusBadRequest}
	// MissingRegionCode represents error message for missingRegionCode
	MissingRegionCode = ResponseError{"invalidRegionCode", "regionCode missing in the input params", http.StatusBadRequest}

	// InvalidCityName represents error message for invalidCityName
	InvalidCityName = ResponseError{"invalidCityName", "invalid cityName provided", http.StatusBadRequest}
	// MissingCityName represents error message for missingCityName
	MissingCityName = ResponseError{"missingCityName", "cityName missing in the input params", http.StatusBadRequest}

	// InvalidBrand represents error message for invalidBrand
	InvalidBrand = ResponseError{"invalidBrand", "invalid brand provided", http.StatusBadRequest}
	// MissingBrand represents error message for missingBrand
	MissingBrand = ResponseError{"missingBrand", "brand missing in the input params", http.StatusBadRequest}

	// InvalidType represents error message for invalidType
	InvalidType = ResponseError{"invalidType", "invalid type provided", http.StatusBadRequest}
	// MissingType represents error message for missingType
	MissingType = ResponseError{"missingType", "type missing in the input params", http.StatusBadRequest}

	// InvalidCategory represents error message for invalidCategory
	InvalidCategory = ResponseError{"invalidCategory", "invalid category provided", http.StatusBadRequest}
	// MissingCategory represents error message for missingCategory
	MissingCategory = ResponseError{"missingCategory", "category missing in the input params", http.StatusBadRequest}

	// InvalidSubCategory represents error message for invalidSubCategory
	InvalidSubCategory = ResponseError{"invalidSubCategory", "invalid subCategory provided", http.StatusBadRequest}
	// MissingSubCategory represents error message for missingSubCategory
	MissingSubCategory = ResponseError{"missingSubCategory", "subCategory missing in the input params", http.StatusBadRequest}

	// InvalidLanguageCodes represents error message for invalidLanguageCodes
	InvalidLanguageCodes = ResponseError{"invalidLanguageCodes", "invalid languageCodes provided", http.StatusBadRequest}
	// MissingLanguageCodes represents error message for missingLanguageCodes
	MissingLanguageCodes = ResponseError{"missingLanguageCodes", "languageCodes missing in the input params", http.StatusBadRequest}

	// MissingActionId represents error message for missingLanguageCodes
	MissingActionId = ResponseError{"missingActionId", "actionId missing in the input params", http.StatusBadRequest}
	// InvalidActionId represents error message for InvalidActionId
	InvalidActionId = ResponseError{"InvalidActionId", "invalid actionId provided", http.StatusBadRequest}

	// MissingKeywords represents error message for MissingKeywords
	MissingKeywords = ResponseError{"missingKeywords", "keywords missing in the request", http.StatusBadRequest}

	// ErrorConfigExists represents error message for configExists
	ErrorConfigExists = ResponseError{"configExists", "Config with given ID already exist", http.StatusBadRequest}
	// ErrorConfigNotFound represents error message for configNotFound
	ErrorConfigNotFound = ResponseError{"configNotFound", "Config with given ID does not exist", http.StatusBadRequest}
	// ErrMissingUserServiceURI is returned when env variable is missing.
	ErrMissingUserServiceURI = errors.New("missingEnvironmentVariable: The value of environment variable 'API_BASE_URL' is missing")
	// ErrInvalidUserServieURI is returned when the environment varible related to url is not properly set.
	ErrInvalidUserServieURI = errors.New("invalidUserServieURI: URI 'API_BASE_URL' must be absolute i.e it should have scheme(eg: http, https etc)")

	// InvalidUserName represents error message for invalidUserName
	InvalidUserName = ResponseError{"invalidUserName", "invalid userName provided", http.StatusBadRequest}
	// InvalidPhoneNumber represents error message for invalidPhoneNumber
	InvalidPhoneNumber = ResponseError{"invalidPhoneNumber", "invalid phoneNumber provided", http.StatusBadRequest}
	// InvalidCustomerNumber represents error message for invalid invalidCustomerPhoneNumber
	InvalidCustomerNumber = ResponseError{"invalidCustomerPhoneNumber", "invalid Customer phoneNumber provided", http.StatusBadRequest}
	// InvalidState	represents error message for invalidState
	InvalidState = ResponseError{"invalidState", "invalid state provided", http.StatusBadRequest}
	// InvalidCity	represents error message for invalidCity
	InvalidCity = ResponseError{"invalidCity", "invalid city provided", http.StatusBadRequest}
	// InvalidPinCode represents error message for invalidPinCode
	InvalidPinCode = ResponseError{"invalidPinCode", "invalid pinCode provided", http.StatusBadRequest}
	// InvalidAddress represents error message for invalidAddress
	InvalidAddress = ResponseError{"invalidAddress", "invalid address provided", http.StatusBadRequest}
	// InvalidAlternatePhoneNumber represents error message for invalidAlternatePhoneNumber
	InvalidAlternatePhoneNumber = ResponseError{"invalidAlternatePhoneNumber", "invalid alternatePhoneNumber provided", http.StatusBadRequest}
	// InvalidAddressName represents error message for invalidAddressName
	InvalidAddressName = ResponseError{"invalidAddressName", "invalid addressName provided", http.StatusBadRequest}
	// InvalidUserAddressPayload represents error message for invalidUserAddressPayload
	InvalidUserAddressPayload = ResponseError{"invalidUserAddressPayload", "invalid userAddressPayload provided", http.StatusBadRequest}
	// InvalidAccessToken represents error message for invalidAccessToken
	InvalidAccessToken = ResponseError{"invalidAccessToken", "invalid accessToken provided", http.StatusUnauthorized}
	// InvalidAPIKey represents error message for invalidAPIKey
	InvalidAPIKey = ResponseError{"invalidAPIKey", "invalid api key provided", http.StatusUnauthorized}
	// InvalidAPIKey represents error message for invalidAPIKey
	MissingAPIKey = ResponseError{"missingAPIKey", "missing api key", http.StatusUnauthorized}
	// InvalidUserAddressID represents error message for invalidUserAddressID
	InvalidUserAddressID = ResponseError{"invalidUserAddressID", "invalid userAddressID provided", http.StatusBadRequest}
	// InvalidAddressType represents error message for nvalidAddressType
	InvalidAddressType = ResponseError{"invalidAddressType", "invalid addressType provided", http.StatusBadRequest}

	// InvalidVendorName represents error message for invalidVendorName
	InvalidVendorName = ResponseError{"invalidVendorname", "invalid vendor name provided", http.StatusBadRequest}

	//MissingVendorName represents error message for missingVendorName
	MissingVendorName = ResponseError{"missingVendorname", "vendor name is missing in the params provided", http.StatusBadRequest}
	// ErrUnexpectedError is returned when an unexpected error occurs.
	ErrUnexpectedError = ResponseError{"unexpectedError", "An unexpected error occurred", http.StatusInternalServerError}
	// InvalidRequestPayload represents error message for invalidRequestPayload
	InvalidRequestPayload = ResponseError{"invalidRequestPayload", "invalid request payload provided", http.StatusBadRequest}
	// MissingProductRequestPayload represents error message for missingProductRequestPayload
	MissingProductRequestPayload = ResponseError{"missingProductRequestPayload", "product request payload is missing", http.StatusBadRequest}
	// ProductRequestPayloadExceedsMaxLength represents error message for productRequestPayloadExceedsMaxLength
	ProductRequestPayloadExceedsMaxLength = ResponseError{"productRequestPayloadExceedsMaxLength", "product request payload exceeds max length", http.StatusBadRequest}
	// InvalidUserRegisterationPayload represents error message for invalidUserRegisterationPayload
	InvalidUserRegisterationPayload = ResponseError{"invalidUserRegisterationPayload", "invalid user registeration payload provided", http.StatusBadRequest}
	// InvalidUserVotesPayload represents error message for invalidUserRegisterationPayload
	InvalidUserVotesPayload = ResponseError{"invalidUserVotesPayload", "invalid user votes payload provided", http.StatusBadRequest}

	// MissingCloudAccountID represents error message for missingCloudAccountID
	MissingCloudAccountID = ResponseError{"missingCloudAccountID", "cloud account id is missing", http.StatusBadRequest}
	// InvalidAddOrRemoveFromWishlistPayload represents error message for invalidAddOrRemoveFromWishlistPayload
	InvalidAddOrRemoveFromWishlistPayload = ResponseError{"invalidAddOrRemoveFromWishlistPayload", "invalid add or remove from wishlist payload provided", http.StatusBadRequest}
	// MissingUserID represents error message for missingUserID
	MissingUserID = ResponseError{"missingUserID", "user id is missing", http.StatusBadRequest}
	// MissingProductID represents error message for missingProductID
	MissingProductID = ResponseError{"missingProductID", "product id is missing", http.StatusBadRequest}
	// MissingVariantID represents error message for missingVariantID
	MissingVariantID = ResponseError{"missingVariantID", "variant id is missing", http.StatusBadRequest}
	// MissingVendorID represents error message for missingVendorID
	MissingVendorID = ResponseError{"missingVendorID", "vendor id is missing", http.StatusBadRequest}
	// MissingQuantity represents error message for missingQuantity
	MissingQuantity = ResponseError{"missingQuantity", "quantity is missing", http.StatusBadRequest}
	// InvalidUserID represents error message for invalidUserID
	InvalidUserID = ResponseError{"invalidUserID", "invalid user id provided", http.StatusBadRequest}
	// InvalidQuantity represents error message for invalidQuantity
	InvalidQuantity = ResponseError{"invalidQuantity", "invalid quantity provided", http.StatusBadRequest}
	// InvalidIncrementCartProductQuantityPayload represents error message for invalidIncrementCartProductQuantityPayload
	InvalidIncrementCartProductQuantityPayload = ResponseError{"invalidIncrementCartProductQuantityPayload", "invalid increment cart product quantity payload provided", http.StatusBadRequest}
	// InvalidDecrementCartProductQuantityPayload represents error message for invalidDecrementCartProductQuantityPayload
	InvalidDecrementCartProductQuantityPayload = ResponseError{"invalidDecrementCartProductQuantityPayload", "invalid decrement cart product quantity payload provided", http.StatusBadRequest}
	// InvalidValueForIncrementBy represents error message for invalidValueForIncrementBy
	InvalidValueForIncrementBy = ResponseError{"invalidValueForIncrementBy", "invalid value for increment by provided", http.StatusBadRequest}
	//  InvalidValueForDecrementBy represents error message for invalidValueForDecrementBy
	InvalidValueForDecrementBy = ResponseError{"invalidValueForDecrementBy", "invalid value for decrement by provided", http.StatusBadRequest}
	// InvalidVendorID represents error message for invalidVendorID
	InvalidVendorID = ResponseError{"invalidVendorID", "invalid vendor id provided", http.StatusBadRequest}
	// InvalidUserDpandaLoginPayload represents error message for invalidUserDpandaLoginPayload
	InvalidUserDpandaLoginPayload = ResponseError{"invalidUserDpandaLoginPayload", "invalid user dpanda login payload provided", http.StatusBadRequest}
	// InvalidDpandaAccessToken represents error message for invalidDpandaAccessToken
	InvalidDpandaAccessToken = ResponseError{"invalidDpandaAccessToken", "invalid dpanda access token provided", http.StatusUnauthorized}
	// InvalidDpandaLoginPayload represents error message for invalidDpandaLoginPayload
	InvalidDpandaLoginPayload = ResponseError{"invalidDpandaLoginPayload", "invalid dpanda login payload provided", http.StatusUnauthorized}
	// InternalServerErrorFromDpanda represents error message for InternalServerErrorFromDpanda
	InternalServerErrorFromDpanda = ResponseError{"internalServerErrorFromDpanda", "invalid dpanda payload", http.StatusInternalServerError}
	// InvalidDpandaPayload represents error message for invalidDpandaPayload
	InvalidDpandaPayload = ResponseError{"invalidDpandaPayload", "invalid dpanda payload", http.StatusBadRequest}
	// InvalidEventsPayload represents error message for invalidEventsPayload
	InvalidVendorEventsPayload = ResponseError{"invalidEventsPayload", "invalid events payload", http.StatusBadRequest}
	//IngestionServerFailed represents error message when ingestion server is down
	IngestionServerFailed = ResponseError{"ingestion server is down", "please wait for sometime and try again later", http.StatusBadRequest}
	// InvalidOrderID represents error message for invalidOrderID
	InvalidOrderID = ResponseError{"invalidOrderID", "invalid order id provided", http.StatusBadRequest}
	// InvalidOrderAckowledgementPayload represents error message for invalidOrderAckowledgementPayload
	InvalidOrderAckowledgementPayload = ResponseError{"invalidOrderAckowledgementPayload", "invalid order acknowledgement payload provided", http.StatusBadRequest}
	// InvalidOrderStatusPayload represents error message for invalidOrderStatusPayload
	InvalidOrderStatusPayload = ResponseError{"invalidOrderStatusPayload", "invalid order status payload provided", http.StatusBadRequest}
	// InvalidUpdateCartPayload represents error message for invalidUpdateCartPayload
	InvalidUpdateCartPayload = ResponseError{"invalidUpdateCartPayload", "invalid update cart payload provided", http.StatusBadRequest}
	// InvalidPayload respresents error message for invalidPayload
	InvalidPayload = ResponseError{"invalidPayload", "invalid payload provided", http.StatusBadRequest}
	// InvalidDeviceID represents error message for invalidDeviceID
	InvalidDeviceIDPayload = ResponseError{"invalidDeviceID", "invalid user registeration payload provided", http.StatusBadRequest}
	// InvalidPhoneNumberOrUUID represents error message for invalidPhoneNumberOrUUID
	InvalidPhoneNumberOrUUID = ResponseError{"invalidPhoneNumberOrUUID", "invalid phoneNumber or UUID provided", http.StatusBadRequest}
	// InvalidTimeSlotFormat represents error message for invalidTimeSlotFormat
	InvalidTimeSlotFormat = ResponseError{"invalidTimeSlotFormat", "invalid time slot format provided", http.StatusBadRequest}
	// MissingPhoneNumberAndUUID represents error message for missing uuid and phone number
	MissingPhoneNumberAndUUID = ResponseError{"missingPhoneNumberAndUUID", "either phone number or uuid should be present in payload", http.StatusBadRequest}
	// MissingSessionID represents error message for missing sessionID
	MissingSessionID = ResponseError{"missingSessionID", "session id is missing in payload", http.StatusBadRequest}
	// MissingSource represents error message for missing source
	MissingSource = ResponseError{"missingSource", "source is missing in payload", http.StatusBadRequest}
	// MissingScreen represents error message for missing action
	MissingScreen = ResponseError{"missingScreen", "screen is missing in payload", http.StatusBadRequest}
	//MissingName represents error message for missing name
	MissingName = ResponseError{"missingName", "name is missing in payload", http.StatusBadRequest}

	//MissingSuperAppId represents error message for missingSuperAppId
	MissingSuperAppId = ResponseError{"missingSuperAppId", "missing super app id", http.StatusBadRequest}
	//InvalidSuperAppId represents error message for invalidSuperAppId
	InvalidSuperAppId = ResponseError{"invalidSuperAppId", "invalid super app id", http.StatusBadRequest}
)

// RedisError represents error code for redis error
const RedisError = "redisError"

// ErrInvalidCountryCode  represents error code for Country not found
var ErrInvalidCountryCode = "invalidCountryCode : Country details not found for ISO Code : "
