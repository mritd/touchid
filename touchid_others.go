// +build !darwin

package touchid

import "errors"

func Authenticate(dType int, reason string) (bool, error) {
	return false, errors.New("touch id auth not support")
}
