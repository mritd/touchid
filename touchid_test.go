package touchid

import "testing"

func TestAuthenticate(t *testing.T) {
	ok, err := Authenticate("Test touch id.")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("touch id authentication result:", ok)
}

func TestAuthenticateWithBiometrics(t *testing.T) {
	ok, err := AuthenticateWithBiometrics("Test touch id(only use biometrics).")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("touch id authentication result:", ok)
}

func TestAuthenticateWithWatch(t *testing.T) {
	ok, err := AuthenticateWithWatch("Test touch id(only use apple watch).")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("touch id authentication result:", ok)
}

func TestAuthenticateWithBiometricsOrWatch(t *testing.T) {
	ok, err := AuthenticateWithBiometricsOrWatch("Test touch id(use biometrics and apple watch).")
	if err != nil {
		t.Fatal(err)
	}
	t.Log("touch id authentication result:", ok)
}
