// +build !darwin

package touchid

import "errors"

func Authenticate(_ string) (bool, error) {
	return false, errors.New("touch id auth not support")
}

func AuthenticateWithBiometrics(_ string) (bool, error) {
	return false, errors.New("touch id auth not support")
}

func AuthenticateWithWatch(_ string) (bool, error) {
	return false, errors.New("touch id auth not support")
}

func AuthenticateWithBiometricsOrWatch(_ string) (bool, error) {
	return false, errors.New("touch id auth not support")
}
