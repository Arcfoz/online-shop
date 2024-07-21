package response

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound       = errors.New("not found")
	ErrUnauthorized   = errors.New("unauthorized")
	ErrForbidenAccess = errors.New("forbiden access")
)

var (
	// auth
	ErrEmailRequired         = errors.New("email is required")
	ErrEmailInvalid          = errors.New("email is invalid")
	ErrPasswordRequired      = errors.New("password is required")
	ErrPasswordInvalidLength = errors.New("password must have minimum 6 character")
	ErrAuthIsNotExist        = errors.New("auth is not exist")
	ErrEmailAlreadyExist     = errors.New("email already used")
	ErrPasswordNotMatch      = errors.New("password not match")

	//product
	ErrProductNameRequired = errors.New("produt name is required")
	ErrProductNameInvalid  = errors.New("product name must have minimum 4 character")
	ErrProductStockInvalid = errors.New("produt stock must be greater than 0")
	ErrProductPriceInvalid = errors.New("produt price must be greater than 0")

	//transaction
	ErrInvalidAmount          = errors.New("invalid amount")
	ErrAmountGreaterThanStock = errors.New("amount greater than stock")
)

type Error struct {
	Message  string
	Code     string
	HttpCode int
}

func NewError(msg string, code string, httpcode int) Error {
	return Error{
		Message:  msg,
		Code:     code,
		HttpCode: httpcode,
	}
}

func (e Error) Error() string {
	return e.Message
}

var (
	ErrorGeneral        = NewError("general error", "99999", http.StatusInternalServerError)
	ErrorBadRequest     = NewError("bad request", "40000", http.StatusBadRequest)
	ErrorNotFound       = NewError(ErrNotFound.Error(), "40400", http.StatusNotFound)
	ErrorUnauthorized   = NewError(ErrUnauthorized.Error(), "40100", http.StatusUnauthorized)
	ErrorForbidenAccess = NewError(ErrForbidenAccess.Error(), "40100", http.StatusForbidden)
)

var (
	// Error bad request
	ErrorEmailRequired          = NewError(ErrEmailRequired.Error(), "40001", http.StatusBadRequest)
	ErrorEmailInvalid           = NewError(ErrEmailInvalid.Error(), "40002", http.StatusBadRequest)
	ErrorPasswordRequired       = NewError(ErrPasswordRequired.Error(), "40003", http.StatusBadRequest)
	ErrorPasswordInvalidLength  = NewError(ErrPasswordInvalidLength.Error(), "40004", http.StatusBadRequest)
	ErrorProductNameRequired    = NewError(ErrProductNameRequired.Error(), "40005", http.StatusBadRequest)
	ErrorProductNameInvalid     = NewError(ErrProductNameInvalid.Error(), "40006", http.StatusBadRequest)
	ErrorProductStockInvalid    = NewError(ErrProductStockInvalid.Error(), "40007", http.StatusBadRequest)
	ErrorProductPriceInvalid    = NewError(ErrProductPriceInvalid.Error(), "40008", http.StatusBadRequest)
	ErrorInvalidAmount          = NewError(ErrInvalidAmount.Error(), "40009", http.StatusBadGateway)
	ErrorAmountGreaterThanStock = NewError(ErrAmountGreaterThanStock.Error(), "40010", http.StatusBadGateway)

	ErrorAuthIsNotExist    = NewError(ErrAuthIsNotExist.Error(), "40401", http.StatusNotFound)
	ErrorEmailAlreadyExist = NewError(ErrEmailAlreadyExist.Error(), "40901", http.StatusConflict)
	ErrorPasswordNotMatch  = NewError(ErrPasswordNotMatch.Error(), "40101", http.StatusUnauthorized)
)

var (
	ErrorMapping = map[string]Error{
		ErrNotFound.Error():              ErrorNotFound,
		ErrEmailRequired.Error():         ErrorEmailRequired,
		ErrEmailInvalid.Error():          ErrorEmailInvalid,
		ErrPasswordRequired.Error():      ErrorPasswordRequired,
		ErrPasswordInvalidLength.Error(): ErrorPasswordInvalidLength,
		ErrAuthIsNotExist.Error():        ErrorAuthIsNotExist,
		ErrEmailAlreadyExist.Error():     ErrorEmailAlreadyExist,
		ErrPasswordNotMatch.Error():      ErrorPasswordNotMatch,
		ErrUnauthorized.Error():          ErrorUnauthorized,
		ErrForbidenAccess.Error():        ErrorForbidenAccess,
	}
)
