package benchmarkmeasure

import (
	"errors"
	"os"
	"testing"
	"time"

	"github.com/pedrodrim/bml-golang/src/model/exception/benchmarkException"
	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
	"github.com/pedrodrim/bml-golang/src/model/timeFormat"
)

func TestBenchmarkMeasure(t *testing.T) {
	t.Run("1. Deverá retornar InvalidParameterException ao chamar 'Start' com 'tag' vazia", func(t *testing.T) {
		instance := NewBenchmarkMeasure()

		err := instance.Start("")

		if err == nil {
			t.Fatal("1. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("1. Esperava *InvalidParameterException, recebeu %T", err)
		}
		if err.Error() != "'tag' é nil" {
			t.Errorf("1. Esperava %q, recebeu %q", "'tag' é nil", err.Error())
		}
	})

	t.Run("2. Deverá retornar InvalidParameterException ao chamar 'End' com 'tag' vazia", func(t *testing.T) {
		instance := NewBenchmarkMeasure()

		err := instance.End("")

		if err == nil {
			t.Fatal("2. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("2. Esperava *InvalidParameterException, recebeu %T", err)
		}
		if err.Error() != "'tag' é nil" {
			t.Errorf("2. Esperava %q, recebeu %q", "'tag' é nil", err.Error())
		}
	})

	t.Run("3. Deverá retornar InvalidParameterException ao chamar 'ResultByTag' com 'tag' vazia", func(t *testing.T) {
		instance := NewBenchmarkMeasure()

		_, err := instance.ResultByTag("", timeformat.Segundos)

		if err == nil {
			t.Fatal("3. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("3. Esperava *InvalidParameterException, recebeu %T", err)
		}
		if err.Error() != "'tag' é nil" {
			t.Errorf("3. Esperava %q, recebeu %q", "'tag' é nil", err.Error())
		}
	})

	t.Run("4. Deverá retornar InvalidParameterException ao chamar 'ResultByTag' com 'format' zerado", func(t *testing.T) {
		instance := NewBenchmarkMeasure()

		_, err := instance.ResultByTag("test", 0)

		if err == nil {
			t.Fatal("4. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("4. Esperava *InvalidParameterException, recebeu %T", err)
		}
		if err.Error() != "'format' é nil" {
			t.Errorf("4. Esperava %q, recebeu %q", "'format' é nil", err.Error())
		}
	})

	t.Run("5. Deverá retornar BenchmarkException ao chamar 'ResultByTag' com tag não iniciada/finalizada", func(t *testing.T) {
		instance := NewBenchmarkMeasure()

		_, err := instance.ResultByTag("test", timeformat.Segundos)

		if err == nil {
			t.Fatal("5. Esperava erro, recebeu nil")
		}
		var target *benchmarkexception.BenchmarkException
		if !errors.As(err, &target) {
			t.Errorf("5. Esperava *BenchmarkException, recebeu %T", err)
		}
		expected := "Não encontrado par 'inicio-fim' de:test"
		if err.Error() != expected {
			t.Errorf("5. Esperava %q, recebeu %q", expected, err.Error())
		}
	})

	t.Run("6. Deverá chamar 'Start' e 'End' corretamente e retornar o tempo de execução", func(t *testing.T) {
		instance := NewBenchmarkMeasure()

		if err := instance.Start("test"); err != nil {
			t.Fatalf("6. Start retornou erro inesperado: %v", err)
		}
		time.Sleep(1 * time.Second)
		if err := instance.End("test"); err != nil {
			t.Fatalf("6. End retornou erro inesperado: %v", err)
		}

		result, err := instance.ResultByTag("test", timeformat.Milissegundos)
		if err != nil {
			t.Fatalf("6. ResultByTag retornou erro inesperado: %v", err)
		}

		// Como Milissegundos = 1 (fator), o resultado deve estar ~1000.
		// Aceita uma folga para variação de agendamento do SO.
		if result < 1000 || result > 2000 {
			t.Errorf("6. Esperava resultado entre 1000 e 2000, recebeu %v", result)
		}
	})

	t.Run("7. Deverá exportar os resultados para um arquivo corretamente", func(t *testing.T) {
		instance := NewBenchmarkMeasure()

		if err := instance.Start("test"); err != nil {
			t.Fatalf("7. Start retornou erro inesperado: %v", err)
		}
		time.Sleep(1 * time.Second)
		if err := instance.End("test"); err != nil {
			t.Fatalf("7. End retornou erro inesperado: %v", err)
		}

		fileName := "benchmark_result.json"
		defer os.Remove(fileName) // limpeza mesmo em caso de falha

		if err := instance.Export(fileName, timeformat.Milissegundos); err != nil {
			t.Fatalf("7. Export retornou erro inesperado: %v", err)
		}

		if _, err := os.Stat(fileName); errors.Is(err, os.ErrNotExist) {
			t.Errorf("7. Esperava que o arquivo %q existisse", fileName)
		}
	})
}