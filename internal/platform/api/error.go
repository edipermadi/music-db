package api

// Error represents api Error
type Error struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

// Error message enumerations
var (
	ErrInternalServer    = Error{Code: 500101, Message: "internal server error"}
	ErrBadQueryParameter = Error{Code: 400101, Message: "bad query parameter"}
	ErrResourceNotFound  = Error{Code: 404101, Message: "resource not found"}
)
