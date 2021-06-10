package utils

import "fmt"

func InternalProcessingErr(err string) error {
	return fmt.Errorf("internal processing error: %s", err)
}
