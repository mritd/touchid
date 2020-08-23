package touchid

import "testing"

func TestAuthenticate(t *testing.T) {
	ok, err := Auth(DeviceTypeAny, "Test touch id.")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("touch id authentication result:", ok)
}

func TestAuthenticateWithWatch(t *testing.T) {
	ok, err := Auth(DeviceTypeWatch, "Test touch id(only use apple watch).")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("touch id authentication result:", ok)
}

func TestAuthenticateWithBiometrics(t *testing.T) {
	ok, err := Auth(DeviceTypeBiometrics, "Test touch id(only use biometrics).")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("touch id authentication result:", ok)
}

func TestAuthenticateWithBiometricsOrWatch(t *testing.T) {
	ok, err := Auth(DeviceTypeBiometricsOrWatch, "Test touch id(use biometrics and apple watch).")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("touch id authentication result:", ok)
}
