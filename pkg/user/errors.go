package user

import "fmt"

func ExistingEmailErr(email string) error {
	return fmt.Errorf("email %s already in use", email)
}

func InternalProcessingErr(err string) error {
	return fmt.Errorf("internal processing error: %s", err)
}
