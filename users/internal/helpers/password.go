package helpers

import "golang.org/x/crypto/bcrypt"

func HashPassword(rawPassword string) (string, error) {
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(rawPassword), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashPassword), nil
}

func CheckPassword(userPassword, validPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(validPassword), []byte(userPassword))
	if err != nil {
		return err
	}
	return nil
}
