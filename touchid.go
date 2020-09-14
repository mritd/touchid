package touchid

import (
	"errors"
	"sync"
	"time"
)

type DeviceType string

const (
	DeviceTypeAny               DeviceType = "any"
	DeviceTypeWatch             DeviceType = "watch"
	DeviceTypeBiometrics        DeviceType = "biometrics"
	DeviceTypeBiometricsOrWatch DeviceType = "biometrics_watch"
)

var serialLock sync.Mutex
var lastAuthTime time.Time

func SerialAuth(dType DeviceType, reason string, timeout time.Duration) (bool, error) {
	serialLock.Lock()
	defer serialLock.Unlock()
	if lastAuthTime.Add(timeout).After(time.Now()) {
		return true, nil
	} else {
		ok, err := Auth(dType, reason)
		if ok && err == nil {
			lastAuthTime = time.Now()
		}
		return ok, err
	}
}

func Auth(dType DeviceType, reason string) (bool, error) {
	switch dType {
	case DeviceTypeAny:
		return Authenticate(1, reason)
	case DeviceTypeWatch:
		return Authenticate(2, reason)
	case DeviceTypeBiometrics:
		return Authenticate(3, reason)
	case DeviceTypeBiometricsOrWatch:
		return Authenticate(4, reason)
	default:
		return false, errors.New("invalid device type")
	}
}
