package user

import "fmt"

func ExistingEmailErr(email string) error {
	return fmt.Errorf("email %s already in use", email)
}
