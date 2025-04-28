package userInfo

import (
	"testing"
)

func TestUserInfo(t *testing.T) {
	tests := []struct {
		name        string
		setup       func() *UserInfo
		assertions  func(t *testing.T, user *UserInfo)
	}{
		{
			name: "1. Deverá instanciar corretamente",
			setup: func() *UserInfo {
				return NewUserInfo("user", "password", 0)
			},
			assertions: func(t *testing.T, user *UserInfo) {
				if user == nil {
					t.Error("expected NewUserInfo to return a non-nil instance")
				}
			},
		},
		{
			name: "2. 'username' deverá ser alterável",
			setup: func() *UserInfo {
				user := NewUserInfo("user", "password", 0)
				user.SetUser("newUser")
				return user
			},
			assertions: func(t *testing.T, user *UserInfo) {
				if user.GetUser() != "newUser" {
					t.Errorf("expected username to be 'newUser', got '%s'", user.GetUser())
				}
			},
		},
		{
			name: "3. 'credit' deverá ser alterável",
			setup: func() *UserInfo {
				user := NewUserInfo("user", "password", 0)
				user.SetCredit(10)
				return user
			},
			assertions: func(t *testing.T, user *UserInfo) {
				if user.GetCredit() != 10 {
					t.Errorf("expected credit to be 10, got %.2f", user.GetCredit())
				}
			},
		},
		{
			name: "4. Método 'setPassword()' deverá ser executado corretamente",
			setup: func() *UserInfo {
				user := NewUserInfo("user", "password", 0)
				user.SetPassword("newPassword")
				return user
			},
			assertions: func(t *testing.T, user *UserInfo) {
				if user.GetPassword() == "" {
					t.Error("expected encrypted password to be non-empty")
				}
			},
		},
		{
			name: "5. 'username' deverá ser obtido corretamente",
			setup: func() *UserInfo {
				return NewUserInfo("expectedUser", "password", 0)
			},
			assertions: func(t *testing.T, user *UserInfo) {
				if user.GetUser() != "expectedUser" {
					t.Errorf("expected username to be 'expectedUser', got '%s'", user.GetUser())
				}
			},
		},
		{
			name: "6. 'credit' deverá ser obtido corretamente",
			setup: func() *UserInfo {
				return NewUserInfo("user", "password", 20)
			},
			assertions: func(t *testing.T, user *UserInfo) {
				if user.GetCredit() != 20 {
					t.Errorf("expected credit to be 20, got %.2f", user.GetCredit())
				}
			},
		},
		{
			name: "7. Método 'getPassword()' deverá retornar valor criptografado",
			setup: func() *UserInfo {
				return NewUserInfo("user", "password", 0)
			},
			assertions: func(t *testing.T, user *UserInfo) {
				encrypted := user.GetPassword()
				if encrypted == "" {
					t.Error("expected encrypted password to be non-empty")
				}
				if len(encrypted) < 4 || encrypted[:4] != "HASH" {
					t.Errorf("expected encrypted password to start with 'HASH', got '%s'", encrypted)
				}
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			user := tc.setup()
			tc.assertions(t, user)
		})
	}
}
