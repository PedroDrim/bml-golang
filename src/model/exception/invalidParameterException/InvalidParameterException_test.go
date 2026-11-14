package invalidparameterexception

import (
	"errors"
	"testing"
)

func TestInvalidParameterException(t *testing.T) {
	t.Run("1. Deverá implementar a interface error", func(t *testing.T) {
		var err error = NewInvalidParameterException("Erro de parâmetro inválido", nil)

		if err == nil {
			t.Fatal("1. Esperava uma instância não nula de error")
		}

		var target *InvalidParameterException
		if !errors.As(err, &target) {
			t.Errorf("1. Esperava que o erro fosse *InvalidParameterException, recebeu %T", err)
		}
	})

	t.Run("2. Deverá ser retornada corretamente com uma mensagem", func(t *testing.T) {
		var err error = NewInvalidParameterException("Erro de parâmetro inválido", nil)

		expected := "Erro de parâmetro inválido"
		if err.Error() != expected {
			t.Errorf("2. Esperava %q, recebeu %q", expected, err.Error())
		}
	})
}