package maxvalueanalysis

import (
	"testing"
	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

func TestMaxValueAnalysis(t *testing.T) {

	t.Run("1. Instanciação - Deverá instanciar corretamente", func(t *testing.T) {
		instance := NewMaxValueAnalysis()
		if instance == nil {
			t.Error("1. Esperava instância não nula de MaxValueAnalysis")
		}
	})

	t.Run("2. Método Analysis() - Deverá retornar o valor máximo corretamente para uma lista válida", func(t *testing.T) {
		mock := NewMaxValueAnalysis()
		list := []userinfo.UserInfo{
			*userinfo.NewUserInfo("ua", "pa", 1),
			*userinfo.NewUserInfo("ub", "pb", 2),
			*userinfo.NewUserInfo("uc", "pc", 3),
		}
		result := mock.Analysis(list)
		expected := 3.0
		if result != expected {
			t.Errorf("2. Esperava %v, recebeu %v", expected, result)
		}
	})
}