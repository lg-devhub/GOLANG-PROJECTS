package calculator

import "testing"

func TestSoma(t *testing.T) {
	resultado := Soma(10, 5)

	if resultado != 15 {
		t.Errorf("Esperado 15, obtido %v", resultado)
	}
}

func TestSubtracao(t *testing.T) {
	resultado := Subtracao(10, 5)

	if resultado != 5 {
		t.Errorf("Esperado 5, obtido %v", resultado)
	}
}

func TestMultiplicacao(t *testing.T) {
	resultado := Multiplicacao(10, 5)

	if resultado != 50 {
		t.Errorf("Esperado 50, obtido %v", resultado)
	}
}

func TestDivisao(t *testing.T) {
	resultado := Divisao(10, 5)

	if resultado != 2 {
		t.Errorf("Esperado 2, obtido %v", resultado)
	}
}