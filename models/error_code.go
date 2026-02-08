package models

type ErrorCode string

const (
	// AUTH
	ErrAuthInvalidCredential ErrorCode = "AUTH_INVALID_CREDENTIAL"
	ErrAuthUnauthorized      ErrorCode = "AUTH_UNAUTHORIZED"
	ErrAuthTokenExpired      ErrorCode = "AUTH_TOKEN_EXPIRED"
	ErrUserExists            ErrorCode = "USER_ALREADY_EXISTS"

	// USER
	ErrUserNotFound ErrorCode = "USER_NOT_FOUND"

	// VALIDATION
	ErrValidation ErrorCode = "VALIDATION_ERROR"

	// INTERNAL
	ErrInternal ErrorCode = "INTERNAL_SERVER_ERROR"

	// PRODUCT
	ErrProductNotFound   ErrorCode = "PRODUCT_NOT_FOUND"
	ErrProductOutOfStock ErrorCode = "PRODUCT_OUT_OF_STOCK"
	ErrProductExists     ErrorCode = "PRODUCT_ALREADY_EXISTS"

	// TRANSACTION
	ErrTransactionNotFound ErrorCode = "TRANSACTION_NOT_FOUND"
	ErrTransactionFailed   ErrorCode = "TRANSACTION_FAILED"

	// CATEGORY
	ErrCategoryNotFound ErrorCode = "CATEGORY_NOT_FOUND"
	ErrCategoryExists   ErrorCode = "CATEGORY_ALREADY_EXISTS"
)
