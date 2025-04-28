package table

import (
	"testing"
)

func TestTable(t *testing.T) {
	t.Run("1. Deverá instanciar corretamente", func(t *testing.T) {
		table := NewTable("../../../data/test.csv")
		if table == nil {
			t.Fatal("expected NewTable to return a non-nil instance")
		}
	})

	t.Run("2. Método 'GetUserInfoList()' deverá ser executado corretamente", func(t *testing.T) {
		table := NewTable("../../../data/test.csv")
		userList := table.GetUserInfoList()
		if len(userList) != 10 {
			t.Fatalf("expected user list to have 10 elements, got '%d'", len(userList))
		}
	})
}
