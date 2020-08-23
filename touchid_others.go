// +build !darwin

package touchid

func Authenticate(reason string) (bool, error) {
	return false, errors.New("touch id auth not support")
}

func AuthenticateWithBiometrics(reason string) (bool, error) {
	return false, errors.New("touch id auth not support")
}

func AuthenticateWithWatch(reason string) (bool, error) {
	return false, errors.New("touch id auth not support")
}

func AuthenticateWithBiometricsOrWatch(reason string) (bool, error) {
	return false, errors.New("touch id auth not support")
}