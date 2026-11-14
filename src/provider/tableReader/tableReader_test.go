package tablereader

import (
	"errors"
	"os"
	"testing"

	"github.com/pedrodrim/bml-golang/src/model/exception/dataReaderException"
	"github.com/pedrodrim/bml-golang/src/model/exception/invalidParameterException"
)

// setupCSV cria um arquivo CSV temporário para os testes e retorna o caminho.
// O chamador deve chamar os.Remove no final (ou usar t.Cleanup).
func setupCSV(t *testing.T) string {
	t.Helper()

	content := "user,password,credit\n" +
		"User_1,Password_1,1\n" +
		"User_2,Password_2,2\n" +
		"User_3,Password_3,3\n" +
		"User_4,Password_4,4\n" +
		"User_5,Password_5,5\n"

	tmp, err := os.CreateTemp("", "test_*.csv")
	if err != nil {
		t.Fatalf("Falha ao criar arquivo temporário: %v", err)
	}
	if _, err := tmp.WriteString(content); err != nil {
		t.Fatalf("Falha ao escrever arquivo temporário: %v", err)
	}
	if err := tmp.Close(); err != nil {
		t.Fatalf("Falha ao fechar arquivo temporário: %v", err)
	}
	t.Cleanup(func() { os.Remove(tmp.Name()) })

	return tmp.Name()
}

func TestTableReader(t *testing.T) {
	fileName := setupCSV(t)

	t.Run("1. Deverá ser instanciável corretamente", func(t *testing.T) {
		reader, err := NewTableReader(fileName)

		if err != nil {
			t.Fatalf("1. Esperava nil, recebeu: %v", err)
		}
		if reader == nil {
			t.Fatal("1. Esperava uma instância não nula")
		}
		if _, ok := interface{}(reader).(*TableReader); !ok {
			t.Errorf("1. Esperava *TableReader, recebeu %T", reader)
		}
	})

	t.Run("2. Deverá retornar o nome do arquivo corretamente", func(t *testing.T) {
		reader, err := NewTableReader(fileName)
		if err != nil {
			t.Fatalf("2. Esperava nil, recebeu: %v", err)
		}

		if reader.GetFileName() != fileName {
			t.Errorf("2. Esperava %q, recebeu %q", fileName, reader.GetFileName())
		}
	})

	t.Run("3. Deverá ler todos os dados corretamente", func(t *testing.T) {
		reader, err := NewTableReader(fileName)
		if err != nil {
			t.Fatalf("3. Esperava nil, recebeu: %v", err)
		}

		result := reader.ReadAll()
		if len(result) <= 0 {
			t.Errorf("3. Esperava lista não vazia, recebeu %d itens", len(result))
		}
		if len(result) != 5 {
			t.Errorf("3. Esperava 5 itens, recebeu %d", len(result))
		}
	})

	t.Run("4. Deverá ler um intervalo específico corretamente", func(t *testing.T) {
		reader, err := NewTableReader(fileName)
		if err != nil {
			t.Fatalf("4. Esperava nil, recebeu: %v", err)
		}

		// Ruby [1..3] = 3 elementos (inclusivo). Go [1:4] = 3 elementos.
		result, err := reader.Read(1, 3)
		if err != nil {
			t.Fatalf("4. Esperava nil, recebeu: %v", err)
		}
		if len(result) != 3 {
			t.Errorf("4. Esperava 3 itens, recebeu %d", len(result))
		}
	})

	t.Run("5. Deverá retornar exceção ao instanciar com nome de arquivo vazio", func(t *testing.T) {
		_, err := NewTableReader("")

		if err == nil {
			t.Fatal("5. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("5. Esperava *InvalidParameterException, recebeu %T", err)
		}
	})

	t.Run("6. Deverá retornar exceção ao tentar ler com startIndex menor que 0", func(t *testing.T) {
		reader, err := NewTableReader(fileName)
		if err != nil {
			t.Fatalf("6. Setup falhou: %v", err)
		}

		_, err = reader.Read(-1, 2)

		if err == nil {
			t.Fatal("6. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("6. Esperava *InvalidParameterException, recebeu %T", err)
		}
	})

	t.Run("7. Deverá retornar exceção ao tentar ler com endIndex menor que 0", func(t *testing.T) {
		reader, err := NewTableReader(fileName)
		if err != nil {
			t.Fatalf("7. Setup falhou: %v", err)
		}

		_, err = reader.Read(1, -2)

		if err == nil {
			t.Fatal("7. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("7. Esperava *InvalidParameterException, recebeu %T", err)
		}
	})

	t.Run("8. Deverá retornar exceção ao tentar ler com startIndex maior ou igual a endIndex", func(t *testing.T) {
		reader, err := NewTableReader(fileName)
		if err != nil {
			t.Fatalf("8. Setup falhou: %v", err)
		}

		_, err = reader.Read(3, 3)

		if err == nil {
			t.Fatal("8. Esperava erro, recebeu nil")
		}
		var target *invalidparameterexception.InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("8. Esperava *InvalidParameterException, recebeu %T", err)
		}
	})

	t.Run("9. Deverá retornar exceção ao tentar ler um arquivo inexistente", func(t *testing.T) {
		_, err := NewTableReader("inexistent_file.csv")

		if err == nil {
			t.Fatal("9. Esperava erro, recebeu nil")
		}
		var target *datareaderexception.DataReaderException
		if !errors.As(err, &target) {
			t.Errorf("9. Esperava *DataReaderException, recebeu %T", err)
		}
	})
}