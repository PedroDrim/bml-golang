package languagesortanalysis

import (
	"errors"
	"testing"

	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
	"github.com/pedrodrim/bml-golang/src/model/tableAnalysis"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

func TestLanguageSortAnalysis(t *testing.T) {
	t.Run("1. Deverá ser instanciável corretamente", func(t *testing.T) {
		instance := NewLanguageSortAnalysis()

		if instance == nil {
			t.Fatal("1. Esperava uma instância não nula")
		}

		if _, ok := interface{}(instance).(*LanguageSortAnalysis); !ok {
			t.Errorf("1. Esperava *LanguageSortAnalysis, recebeu %T", instance)
		}
	})

	t.Run("2. Deverá implementar a interface TableAnalysis", func(t *testing.T) {
		var iface tableanalysis.TableAnalysis[[]userinfo.UserInfo] = NewLanguageSortAnalysis()

		if iface == nil {
			t.Fatal("2. Esperava que LanguageSortAnalysis implementasse TableAnalysis")
		}
	})

	t.Run("3. Deverá retornar InvalidParameterException para lista nil", func(t *testing.T) {
		instance := NewLanguageSortAnalysis()

		_, err := instance.Analysis(nil)

		if err == nil {
			t.Fatal("3. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("3. Esperava *InvalidParameterException, recebeu %T", err)
		}
		if err.Error() != "'userInfoList' é nil ou vazio" {
			t.Errorf("3. Esperava %q, recebeu %q", "'userInfoList' é nil ou vazio", err.Error())
		}
	})

	t.Run("4. Deverá retornar InvalidParameterException para lista vazia", func(t *testing.T) {
		instance := NewLanguageSortAnalysis()

		_, err := instance.Analysis([]userinfo.UserInfo{})

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

	t.Run("5. Deverá ordenar a lista por crédito decrescente", func(t *testing.T) {
		instance := NewLanguageSortAnalysis()

		list := []userinfo.UserInfo{
			*userinfo.NewUserInfo("ua", "pa", 1),
			*userinfo.NewUserInfo("ub", "pb", 3),
			*userinfo.NewUserInfo("uc", "pc", 2),
		}

		result, err := instance.Analysis(list)

		if err != nil {
			t.Fatalf("5. Esperava nil, recebeu: %v", err)
		}
		if len(result) != 3 {
			t.Fatalf("5. Esperava 3 itens, recebeu %d", len(result))
		}
		if result[0].GetUser() != "ub" {
			t.Errorf("5. Esperava 'ub' na posição 0, recebeu %q", result[0].GetUser())
		}
		if result[1].GetUser() != "uc" {
			t.Errorf("5. Esperava 'uc' na posição 1, recebeu %q", result[1].GetUser())
		}
		if result[2].GetUser() != "ua" {
			t.Errorf("5. Esperava 'ua' na posição 2, recebeu %q", result[2].GetUser())
		}
	})

	t.Run("6. Deverá retornar lista com único elemento sem erro", func(t *testing.T) {
		instance := NewLanguageSortAnalysis()

		list := []userinfo.UserInfo{
			*userinfo.NewUserInfo("ua", "pa", 5),
		}

		result, err := instance.Analysis(list)

		if err != nil {
			t.Fatalf("6. Esperava nil, recebeu: %v", err)
		}
		if len(result) != 1 {
			t.Fatalf("6. Esperava 1 item, recebeu %d", len(result))
		}
		if result[0].GetUser() != "ua" {
			t.Errorf("6. Esperava 'ua', recebeu %q", result[0].GetUser())
		}
	})

	t.Run("7. Não deverá mutar a lista original do chamador", func(t *testing.T) {
		instance := NewLanguageSortAnalysis()

		list := []userinfo.UserInfo{
			*userinfo.NewUserInfo("ua", "pa", 1),
			*userinfo.NewUserInfo("ub", "pb", 3),
		}

		_, err := instance.Analysis(list)
		if err != nil {
			t.Fatalf("7. Esperava nil, recebeu: %v", err)
		}

		if list[0].GetUser() != "ua" {
			t.Errorf("7. Lista original foi mutada; esperava 'ua' na posição 0, recebeu %q", list[0].GetUser())
		}
	})
}