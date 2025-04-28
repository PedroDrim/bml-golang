package userinfo

import (
	"testing"
)

func TestUserInfo(t *testing.T) {
	t.Run("1. NewUserInfo deverá instanciar corretamente", func(t *testing.T) {
		user := NewUserInfo("user", "password", 0)
		if user == nil {
			t.Fatal("expected NewUserInfo to return a non-nil instance")
		}
	})

	t.Run("2. SetUser deverá alterar o usuário corretamente", func(t *testing.T) {
		user := NewUserInfo("user", "password", 0)
		user.SetUser("newUser")
		if user.GetUser() != "newUser" {
			t.Fatalf("expected username to be 'newUser', got '%s'", user.GetUser())
		}
	})

	t.Run("3. SetCredit deverá alterar o crédito corretamente", func(t *testing.T) {
		user := NewUserInfo("user", "password", 0)
		user.SetCredit(10)
		if user.GetCredit() != 10 {
			t.Fatalf("expected credit to be 10, got %.2f", user.GetCredit())
		}
	})

	t.Run("4. SetPassword deverá alterar a senha corretamente", func(t *testing.T) {
		user := NewUserInfo("user", "password", 0)
		user.SetPassword("newPassword")
		if user.GetPassword() == "" {
			t.Fatal("expected encrypted password to be non-empty")
		}
	})

	t.Run("5. GetUser deverá retornar o usuário corretamente", func(t *testing.T) {
		user := NewUserInfo("expectedUser", "password", 0)
		if user.GetUser() != "expectedUser" {
			t.Fatalf("expected username to be 'expectedUser', got '%s'", user.GetUser())
		}
	})

	t.Run("6. GetCredit deverá retornar o crédito corretamente", func(t *testing.T) {
		user := NewUserInfo("user", "password", 20)
		if user.GetCredit() != 20 {
			t.Fatalf("expected credit to be 20, got %.2f", user.GetCredit())
		}
	})

	t.Run("7. GetPassword deverá retornar valor criptografado corretamente", func(t *testing.T) {
		user := NewUserInfo("user", "password", 0)
		encrypted := user.GetPassword()
		if encrypted == "" {
			t.Fatal("expected encrypted password to be non-empty")
		}
		if len(encrypted) < 4 || encrypted[:4] != "HASH" {
			t.Fatalf("expected encrypted password to start with 'HASH', got '%s'", encrypted)
		}
	})
}
