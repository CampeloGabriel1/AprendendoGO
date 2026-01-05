package main

import "testing"


func TestSoma(t *testing.T) {
	teste := soma (3,2,1)
	resultado := 6

	if teste != resultado {
		t.Errorf("Resultado incorreto: esperado %d, obtido %d", resultado, teste)
	}
}

func TestMultiplica(t *testing.T) {
	teste := multiplica (4,5)
	resultado := 20

	if teste != resultado {
		t.Errorf("Resultado incorreto: esperado %d, obtido %d", resultado, teste)
	}
}

func TestSubtrai(t *testing.T) {
	teste := subtrai (10,4)
	resultado := 6
	if teste != resultado {
		t.Errorf("Resultado incorreto: esperado %d, obtido %d", resultado, teste)
	}
}

func TestDivide(t *testing.T) {
	teste := divide (20,5)
	resultado := 4

	if teste != resultado {
		t.Errorf("Resultado incorreto: esperado %d, obtido %d", resultado, teste)
	}
}