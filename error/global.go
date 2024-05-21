package error

import (
	"github.com/yusologia/go-response"
	"net/http"
)

func ErrLogiaUnauthenticated(internalMsg string) {
	logiares.Error(http.StatusUnauthorized, "Your token may have expired. Please re-login!!", internalMsg, nil)
}

func ErrLogiaBadRequest(internalMsg string) {
	logiares.Error(http.StatusBadRequest, "Bad request!", internalMsg, nil)
}

func ErrLogiaPayloadVeryLarge(internalMsg string) {
	logiares.Error(http.StatusRequestEntityTooLarge, "Your request entity very large!", internalMsg, nil)
}

func ErrLogiaValidation(attributes []interface{}) {
	logiares.Error(http.StatusBadRequest, "Missing Required Parameter", "", attributes)
}

func ErrLogiaNotFound(internalMsg string) {
	logiares.Error(http.StatusNotFound, "Data not found", internalMsg, nil)
}

func ErrXtremePermission() {
	logiares.Error(http.StatusForbidden, "Permission restricted", "", nil)
}
