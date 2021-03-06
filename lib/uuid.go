package lib

import (
	"fmt"

	"github.com/rogpeppe/fastuuid"
)

// GenerateUUID generated random uuid
func GenerateUUID() string {
	uuidGenerator := fastuuid.MustNewGenerator()
	uuid := uuidGenerator.Next()
	return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
