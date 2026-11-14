package benchmarkoutput

import (
	"testing"

	"github.com/pedrodrim/bml-golang/src/model/timeFormat"
)

// Mock da interface BenchmarkOutput
type MockBenchmarkOutput struct{}

// Garante em tempo de compilação que MockBenchmarkOutput implementa BenchmarkOutput.
var _ BenchmarkOutput = (*MockBenchmarkOutput)(nil)

func (m *MockBenchmarkOutput) Start(tag string) error {
	return nil
}

func (m *MockBenchmarkOutput) End(tag string) error {
	return nil
}

func (m *MockBenchmarkOutput) ResultByTag(tag string, format timeformat.TimeFormat) (float64, error) {
	return 0, nil
}

func (m *MockBenchmarkOutput) Result(format timeformat.TimeFormat) (map[string]float64, error) {
	return map[string]float64{}, nil
}

func (m *MockBenchmarkOutput) Export(fileName string, format timeformat.TimeFormat) error {
	return nil
}

func TestBenchmarkOutput(t *testing.T) {
	t.Run("1. Deverá ser instanciável corretamente", func(t *testing.T) {
		var mockInstance BenchmarkOutput = &MockBenchmarkOutput{}

		if mockInstance == nil {
			t.Fatal("1. Esperava uma instância não nula")
		}

		if _, ok := mockInstance.(*MockBenchmarkOutput); !ok {
			t.Errorf("1. Esperava *MockBenchmarkOutput, recebeu %T", mockInstance)
		}
	})

	t.Run("2. Método 'Start' deverá ser executado corretamente", func(t *testing.T) {
		mockInstance := &MockBenchmarkOutput{}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("2. Não esperava panic, recebeu: %v", r)
			}
		}()

		err := mockInstance.Start("test")
		if err != nil {
			t.Errorf("2. Esperava nil, recebeu: %v", err)
		}
	})

	t.Run("3. Método 'End' deverá ser executado corretamente", func(t *testing.T) {
		mockInstance := &MockBenchmarkOutput{}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("3. Não esperava panic, recebeu: %v", r)
			}
		}()

		err := mockInstance.End("test")
		if err != nil {
			t.Errorf("3. Esperava nil, recebeu: %v", err)
		}
	})

	t.Run("4. Método 'ResultByTag' deverá ser executado corretamente", func(t *testing.T) {
		mockInstance := &MockBenchmarkOutput{}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("4. Não esperava panic, recebeu: %v", r)
			}
		}()

		result, err := mockInstance.ResultByTag("test", timeformat.Segundos)
		if err != nil {
			t.Errorf("4. Esperava nil, recebeu: %v", err)
		}
		if result != 0 {
			t.Errorf("4. Esperava 0, recebeu %v", result)
		}
	})

	t.Run("5. Método 'Result' deverá ser executado corretamente", func(t *testing.T) {
		mockInstance := &MockBenchmarkOutput{}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("5. Não esperava panic, recebeu: %v", r)
			}
		}()

		result, err := mockInstance.Result(timeformat.Segundos)
		if err != nil {
			t.Errorf("5. Esperava nil, recebeu: %v", err)
		}
		if result == nil {
			t.Error("5. Esperava mapa vazio, recebeu nil")
		}
		if len(result) != 0 {
			t.Errorf("5. Esperava mapa vazio, recebeu %v", result)
		}
	})

	t.Run("6. Método 'Export' deverá ser executado corretamente", func(t *testing.T) {
		mockInstance := &MockBenchmarkOutput{}

		defer func() {
			if r := recover(); r != nil {
				t.Errorf("6. Não esperava panic, recebeu: %v", r)
			}
		}()

		err := mockInstance.Export("output.json", timeformat.Segundos)
		if err != nil {
			t.Errorf("6. Esperava nil, recebeu: %v", err)
		}
	})
}