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
		return Authenticate(0, reason)
	case DeviceTypeWatch:
		return Authenticate(1, reason)
	case DeviceTypeBiometrics:
		return Authenticate(2, reason)
	case DeviceTypeBiometricsOrWatch:
		return Authenticate(3, reason)
	default:
		return false, errors.New("invalid device type")
	}
}
