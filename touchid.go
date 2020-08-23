package touchid

import "errors"

type DeviceType string

const (
	DeviceTypeAny               DeviceType = "any"
	DeviceTypeWatch             DeviceType = "watch"
	DeviceTypeBiometrics        DeviceType = "biometrics"
	DeviceTypeBiometricsOrWatch DeviceType = "biometrics_watch"
)

func Auth(dType DeviceType, reason string) (bool, error) {
	switch dType {
	case DeviceTypeAny:
		return Authenticate(reason)
	case DeviceTypeWatch:
		return AuthenticateWithWatch(reason)
	case DeviceTypeBiometrics:
		return AuthenticateWithBiometrics(reason)
	case DeviceTypeBiometricsOrWatch:
		return AuthenticateWithBiometricsOrWatch(reason)
	default:
		return false, errors.New("invalid device type")
	}
}
