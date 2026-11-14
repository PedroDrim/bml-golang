package summaryanalysis

import (
	"errors"
	"testing"

	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

func TestSummaryAnalysis(t *testing.T) {
	userInfoList := []userinfo.UserInfo{
		*userinfo.NewUserInfo("User_1", "Password_1", 1),
		*userinfo.NewUserInfo("User_2", "Password_2", 2),
		*userinfo.NewUserInfo("User_3", "Password_3", 3),
		*userinfo.NewUserInfo("User_4", "Password_4", 4),
		*userinfo.NewUserInfo("User_5", "Password_5", 5),
	}

	t.Run("1. Deve retornar o valor mínimo correto", func(t *testing.T) {
		instance := NewSummaryAnalysis()

		result, err := instance.Analysis(userInfoList)

		if err != nil {
			t.Fatalf("1. Esperava nil, recebeu: %v", err)
		}
		if result.Min != 1 {
			t.Errorf("1. Esperava 1, recebeu %v", result.Min)
		}
	})

	t.Run("2. Deve retornar o valor máximo correto", func(t *testing.T) {
		instance := NewSummaryAnalysis()

		result, err := instance.Analysis(userInfoList)

		if err != nil {
			t.Fatalf("2. Esperava nil, recebeu: %v", err)
		}
		if result.Max != 5 {
			t.Errorf("2. Esperava 5, recebeu %v", result.Max)
		}
	})

	t.Run("3. Deve retornar a média correta", func(t *testing.T) {
		instance := NewSummaryAnalysis()

		result, err := instance.Analysis(userInfoList)

		if err != nil {
			t.Fatalf("3. Esperava nil, recebeu: %v", err)
		}
		expected := 3.0
		if result.Mean != expected {
			t.Errorf("3. Esperava %v, recebeu %v", expected, result.Mean)
		}
	})

	t.Run("4. Deve retornar exceção quando 'userInfoList' for nil", func(t *testing.T) {
		instance := NewSummaryAnalysis()

		_, err := instance.Analysis(nil)

		if err == nil {
			t.Fatal("4. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("4. Esperava *InvalidParameterException, recebeu %T", err)
		}
		if err.Error() != "'userInfoList' é nil ou vazio" {
			t.Errorf("4. Esperava %q, recebeu %q", "'userInfoList' é nil ou vazio", err.Error())
		}
	})

	t.Run("5. Deve retornar exceção quando 'userInfoList' for vazio", func(t *testing.T) {
		instance := NewSummaryAnalysis()

		_, err := instance.Analysis([]userinfo.UserInfo{})

		if err == nil {
			t.Fatal("5. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("5. Esperava *InvalidParameterException, recebeu %T", err)
		}
		if err.Error() != "'userInfoList' é nil ou vazio" {
			t.Errorf("5. Esperava %q, recebeu %q", "'userInfoList' é nil ou vazio", err.Error())
		}
	})
}