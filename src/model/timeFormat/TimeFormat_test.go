package timeformat

import "testing"

func TestTimeFormat(t *testing.T) {
	t.Run("1. Deverá conter a constante SEGUNDOS corretamente", func(t *testing.T) {
		if Segundos != 0.001 {
			t.Errorf("1. Esperava 0.001, recebeu %v", Segundos)
		}
	})

	t.Run("2. Deverá conter a constante MILLISEGUNDOS corretamente", func(t *testing.T) {
		if Milissegundos != 1 {
			t.Errorf("2. Esperava 1, recebeu %v", Milissegundos)
		}
	})

	t.Run("3. Deverá conter a constante NANOSSEGUNDOS corretamente", func(t *testing.T) {
		if Nanossegundos != 1000 {
			t.Errorf("3. Esperava 1000, recebeu %v", Nanossegundos)
		}
	})
}