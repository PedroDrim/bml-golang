package userInfo

import (
	"bytes"
)

// Struct responsavel por representar as informações de um usuário
type UserInfo struct {
	user string
	password string
}

// NewUserInfo cria uma nova instância de UserInfo
// @param user Nome do usuario
// @param password Senha do usuario
// @return Retorna um ponteiro para UserInfo
func NewUserInfo(user string, password string) *UserInfo {
	return &UserInfo{
		user: user,
		password: password,
	}
}

// SetUser atualiza o nome do usuário
// @param user Nome do usuario
func (u *UserInfo) SetUser(user string) {
	u.user = user
}

// SetPassword atualiza a senha do usuário
// @param password Senha do usuario
func (u *UserInfo) SetPassword(password string) {
	u.password = password
}

// GetUser retorna o nome do usuário
// @return nome do usuario
func (u *UserInfo) GetUser() string {
	return u.user
}

// GetPassword retorna a senha do usuário criptografada
// @return senha do usuario criptografada
func (u *UserInfo) GetPassword() string {
	return cryptPassword(u.password)
}

// cryptPassword é uma função interna para criptografar a senha
// @param senha a ser criptografada
// @return senha criptografada
func cryptPassword(password string) string {
	runes := []rune(password)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	var buffer bytes.Buffer
	buffer.WriteString("HASH")
	buffer.WriteString(string(runes))
	buffer.WriteString("000")

	return buffer.String()
}
