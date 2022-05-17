package securit

import (
	"golang.org/x/crypto/bcrypt"
)

//HashPassword receber a senha string e convert em um hash de bytes.
func HashPassword(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

//VerificarSenha compara as strings da senha e o hash armazenado no banco de dados
func VerificarSenha(passwordHash, passwordString string) error {
	return bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(passwordString))
}

func Teste(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
