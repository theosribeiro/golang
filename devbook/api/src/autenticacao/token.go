package autenticacao

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

//retorna um token assinado com as permissoes do usuario
func CriarToken(usuarioID uint64) (string, error) {
	//permissoes
	permissoes := jwt.MapClaims{}
	permissoes["authorized"] = true
	permissoes["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissoes["usuarioId"] = usuarioID

	//chave secret
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissoes)
	return token.SignedString([]byte("Secret"))

}
