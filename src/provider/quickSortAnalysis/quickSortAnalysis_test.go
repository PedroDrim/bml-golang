package quicksortanalysis

import (
	"errors"
	"testing"

	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

func TestQuickSortAnalysis(t *testing.T) {
	userInfoList := []userinfo.UserInfo{
		*userinfo.NewUserInfo("User_1", "Password_1", 1),
		*userinfo.NewUserInfo("User_2", "Password_2", 5),
		*userinfo.NewUserInfo("User_3", "Password_3", 3),
		*userinfo.NewUserInfo("User_4", "Password_4", 4),
		*userinfo.NewUserInfo("User_5", "Password_5", 2),
	}

	t.Run("1. Deve retornar exceção quando 'userInfoList' for nil", func(t *testing.T) {
		instance := NewQuickSortAnalysis()

		_, err := instance.Analysis(nil)

		if err == nil {
			t.Fatal("1. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("1. Esperava *InvalidParameterException, recebeu %T", err)
		}
		if err.Error() != "'userInfoList' é nil ou vazio" {
			t.Errorf("1. Esperava %q, recebeu %q", "'userInfoList' é nil ou vazio", err.Error())
		}
	})

	t.Run("2. Deve retornar exceção quando 'userInfoList' for vazio", func(t *testing.T) {
		instance := NewQuickSortAnalysis()

		_, err := instance.Analysis([]userinfo.UserInfo{})

		if err == nil {
			t.Fatal("2. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("2. Esperava *InvalidParameterException, recebeu %T", err)
		}
		if err.Error() != "'userInfoList' é nil ou vazio" {
			t.Errorf("2. Esperava %q, recebeu %q", "'userInfoList' é nil ou vazio", err.Error())
		}
	})

	t.Run("3. A lista deve estar ordenada em ordem decrescente", func(t *testing.T) {
		instance := NewQuickSortAnalysis()

		sortedList, err := instance.Analysis(userInfoList)

		if err != nil {
			t.Fatalf("3. Esperava nil, recebeu: %v", err)
		}

		expected := []float64{5, 4, 3, 2, 1}
		if len(sortedList) != len(expected) {
			t.Fatalf("3. Esperava %d itens, recebeu %d", len(expected), len(sortedList))
		}
		for i, want := range expected {
			if sortedList[i].GetCredit() != want {
				t.Errorf("3. Posição %d: esperava %v, recebeu %v", i, want, sortedList[i].GetCredit())
			}
		}
	})
}