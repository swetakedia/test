package compliance

import (
	"github.com/asaskevich/govalidator"
	"github.com/test/go/address"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
	govalidator.CustomTypeTagMap.Set("test_address", govalidator.CustomTypeValidator(isTestAddress))
}

func isTestAddress(i interface{}, context interface{}) bool {
	addr, ok := i.(string)

	if !ok {
		return false
	}

	_, _, err := address.Split(addr)

	if err == nil {
		return true
	}

	return false
}
