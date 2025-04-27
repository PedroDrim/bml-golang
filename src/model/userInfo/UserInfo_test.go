// model/userinfo_test.go
package userInfo

import "testing"

func TestNewUserInfo(t *testing.T) {
	u := NewUserInfo("admin", "12345")

	if u.GetUser() != "admin" {
		t.Errorf("Esperado 'admin', mas obteve '%s'", u.GetUser())
	}

	expectedEncrypted := "HASH54321000" // Esperado se sua lógica de criptografia inverter a senha e adicionar prefixo/sufixo
	if u.GetPassword() != expectedEncrypted {
		t.Errorf("Senha encriptada incorreta. Esperado '%s', obteve '%s'", expectedEncrypted, u.GetPassword())
	}
}

func TestSetUser(t *testing.T) {
	u := NewUserInfo("admin", "12345")
	u.SetUser("root")

	if u.GetUser() != "root" {
		t.Errorf("SetUser falhou. Esperado 'root', obteve '%s'", u.GetUser())
	}
}

func TestSetPassword(t *testing.T) {
	u := NewUserInfo("admin", "12345")
	u.SetPassword("newpass")

	expectedEncrypted := "HASHssapwen000"
	if u.GetPassword() != expectedEncrypted {
		t.Errorf("SetPassword falhou. Esperado '%s', obteve '%s'", expectedEncrypted, u.GetPassword())
	}
}
