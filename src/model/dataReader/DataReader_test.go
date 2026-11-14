package datareader

import (
	"errors"
	"testing"

	"github.com/pedrodrim/bml-golang/src/model/userInfo"
)

// Mock da interface DataReader
type MockDataReader struct {
	opened     bool
	openErr    error
	readAllErr error
	readErr    error
}

var _ DataReader = (*MockDataReader)(nil)

func (m *MockDataReader) Open() error {
	if m.openErr != nil {
		return m.openErr
	}
	m.opened = true
	return nil
}

func (m *MockDataReader) ReadAll() ([]userinfo.UserInfo, error) {
	if m.readAllErr != nil {
		return nil, m.readAllErr
	}
	return []userinfo.UserInfo{
		*userinfo.NewUserInfo("ua", "pa", 1),
		*userinfo.NewUserInfo("ub", "pb", 2),
	}, nil
}

func (m *MockDataReader) Read(startIndex int, endIndex int) ([]userinfo.UserInfo, error) {
	if m.readErr != nil {
		return nil, m.readErr
	}

	all, err := m.ReadAll()
	if err != nil {
		return nil, err
	}

	if startIndex < 0 {
		startIndex = 0
	}
	if endIndex > len(all) {
		endIndex = len(all)
	}
	if startIndex >= endIndex {
		return []userinfo.UserInfo{}, nil
	}
	return all[startIndex:endIndex], nil
}

func TestDataReader(t *testing.T) {
	t.Run("1. Deverá ser instanciável corretamente", func(t *testing.T) {
		var mockInstance DataReader = &MockDataReader{}

		if mockInstance == nil {
			t.Fatal("1. Esperava uma instância não nula")
		}

		if _, ok := mockInstance.(*MockDataReader); !ok {
			t.Errorf("1. Esperava *MockDataReader, recebeu %T", mockInstance)
		}
	})

	t.Run("2. Método 'Open' deverá ser executado sem erro", func(t *testing.T) {
		mockInstance := &MockDataReader{}

		err := mockInstance.Open()

		if err != nil {
			t.Fatalf("2. Esperava nil, recebeu: %v", err)
		}
		if !mockInstance.opened {
			t.Errorf("2. Esperava que 'opened' fosse true após Open()")
		}
	})

	t.Run("3. Método 'Open' deverá propagar erro quando configurado", func(t *testing.T) {
		expected := errors.New("falha ao abrir arquivo")
		mockInstance := &MockDataReader{openErr: expected}

		err := mockInstance.Open()

		if !errors.Is(err, expected) {
			t.Errorf("3. Esperava %v, recebeu %v", expected, err)
		}
		if mockInstance.opened {
			t.Errorf("3. Não deveria marcar 'opened' quando há erro")
		}
	})

	t.Run("4. Método 'ReadAll' deverá retornar todos os dados", func(t *testing.T) {
		mockInstance := &MockDataReader{}

		result, err := mockInstance.ReadAll()

		if err != nil {
			t.Fatalf("4. Esperava nil, recebeu: %v", err)
		}
		if len(result) != 2 {
			t.Errorf("4. Esperava 2 itens, recebeu %d", len(result))
		}
	})

	t.Run("5. Método 'ReadAll' deverá propagar erro quando configurado", func(t *testing.T) {
		expected := errors.New("falha ao ler arquivo")
		mockInstance := &MockDataReader{readAllErr: expected}

		result, err := mockInstance.ReadAll()

		if !errors.Is(err, expected) {
			t.Errorf("5. Esperava %v, recebeu %v", expected, err)
		}
		if result != nil {
			t.Errorf("5. Esperava nil, recebeu %v", result)
		}
	})

	t.Run("6. Método 'Read' deverá retornar apenas o intervalo solicitado", func(t *testing.T) {
		mockInstance := &MockDataReader{}

		result, err := mockInstance.Read(0, 1)

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

	t.Run("7. Método 'Read' deverá retornar lista vazia para intervalo inválido", func(t *testing.T) {
		mockInstance := &MockDataReader{}

		result, err := mockInstance.Read(5, 3)

		if err != nil {
			t.Fatalf("7. Esperava nil, recebeu: %v", err)
		}
		if len(result) != 0 {
			t.Errorf("7. Esperava lista vazia, recebeu %v", result)
		}
	})

	t.Run("8. Método 'Read' deverá propagar erro quando configurado", func(t *testing.T) {
		expected := errors.New("falha ao ler intervalo")
		mockInstance := &MockDataReader{readErr: expected}

		result, err := mockInstance.Read(0, 1)

		if !errors.Is(err, expected) {
			t.Errorf("8. Esperava %v, recebeu %v", expected, err)
		}
		if result != nil {
			t.Errorf("8. Esperava nil, recebeu %v", result)
		}
	})
}