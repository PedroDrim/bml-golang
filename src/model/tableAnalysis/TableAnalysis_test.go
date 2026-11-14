package tableanalysis

import (
	"testing"

	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Mock da interface TableAnalysis
type MockTableAnalysis struct{}

// Garante em tempo de compilação que MockTableAnalysis implementa TableAnalysis
var _ TableAnalysis[[]userinfo.UserInfo] = (*MockTableAnalysis)(nil)

func (m *MockTableAnalysis) Analysis(userInfoList []userinfo.UserInfo) ([]userinfo.UserInfo, error) {
	return []userinfo.UserInfo{}, nil
}

func TestTableAnalysis(t *testing.T) {
	t.Run("1. Deverá ser instanciável corretamente", func(t *testing.T) {
		var mockInstance TableAnalysis[[]userinfo.UserInfo] = &MockTableAnalysis{}

		if mockInstance == nil {
			t.Fatal("1. Esperava uma instância não nula")
		}

		if _, ok := mockInstance.(*MockTableAnalysis); !ok {
			t.Errorf("1. Esperava *MockTableAnalysis, recebeu %T", mockInstance)
		}
	})

	t.Run("2. Método 'analysis' deverá ser executado corretamente", func(t *testing.T) {
		mockInstance := &MockTableAnalysis{}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("2. Não esperava panic, recebeu: %v", r)
			}
		}()

		_, err := mockInstance.Analysis([]userinfo.UserInfo{})
		if err != nil {
			t.Errorf("2. Esperava nil, recebeu: %v", err)
		}
	})

	t.Run("3. Método 'analysis' deverá retornar uma lista vazia", func(t *testing.T) {
		mockInstance := &MockTableAnalysis{}

		result, err := mockInstance.Analysis([]userinfo.UserInfo{})

		if err != nil {
			t.Fatalf("3. Esperava nil, recebeu: %v", err)
		}
		if result == nil {
			t.Fatal("3. Esperava slice vazio, recebeu nil")
		}
		if len(result) != 0 {
			t.Errorf("3. Esperava lista vazia, recebeu %v", result)
		}
	})
}