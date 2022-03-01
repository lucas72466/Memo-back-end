package public

import "golang.org/x/crypto/bcrypt"

func GenerateHashedPassword(plainPassword string) (string, error) {
	bytePassword := []byte(plainPassword)

	hashedPassword, err := bcrypt.GenerateFromPassword(bytePassword, bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}

	return string(hashedPassword), nil
}

func ComparePasswords(plainPassword, hashedPassword string) (bool, error) {
	bytePlainPassword, byteHashedPassword := []byte(plainPassword), []byte(hashedPassword)

	if err := bcrypt.CompareHashAndPassword(byteHashedPassword, bytePlainPassword); err != nil {
		return false, err
	}

	return true, nil
}
