package utils

import (
	"fmt"
	"log"
)

// ValidateToken - func for validate token API.
func ValidateToken(authID, auth, key string) (bool, error) {
	if authID == "" || auth == "" || key == "" {
		return false, fmt.Errorf("authID/auth/key nil")
	}

	tokens := gettokenOrgs()

	token, ok := tokens[authID]
	if !ok {
		log.Println("tokenID not match")
		return false, nil
	}

	if token != auth {
		log.Println("token not match")
		return false, nil
	}

	return true, nil
}
