package poemas_test

import (
	"testing"

	"github.com/nportas/alfonsina-bot/palabras"
	"github.com/nportas/alfonsina-bot/poemas"
)

func TestNoDebeGenerarPoesiaConVersosVacios(t *testing.T) {

	// Inicialización
	predictor := palabras.NewPredictorMock()
	predictor.ResponderConFrase(" ")
	poemario := poemas.NewPoemario(predictor)

	// Operación
	poesia := poemario.GenerarPoesiaAPartirDe("amor", "1", "3")

	// Validación
	if len(poesia.Estrofas) != 0 {
		t.Logf("No se esperaba una poesía pero se generó '%v'", poesia)
		t.FailNow()
	}

}

func TestDebeGenerarVersoConPalabraFinaleDeMasDe3Letras(t *testing.T) {

	// Inicialización
	predictor := palabras.NewPredictorMock()
	predictor.ResponderConFrase("fuego en tu sombra todo lo real me")
	poemario := poemas.NewPoemario(predictor)

	// Operación
	poesia := poemario.GenerarPoesiaAPartirDe("fuego", "1", "3")

	// Validación
	palabrasDelVerso := poesia.Estrofas[0].Versos[0].Palabras
	if palabrasDelVerso[len(palabrasDelVerso)-1] == "me" {
		t.Logf("Se esperaba un verso sin conectores como palabra final pero se generó '%s'", poesia.Estrofas[0].Versos[0].Palabras)
		t.FailNow()
	}

}

func TestNoDebeGenerarVersoConPalabraFinaleDe1Letra(t *testing.T) {

	// Inicialización
	predictor := palabras.NewPredictorMock()
	predictor.ResponderConFrase("beso viene a")
	poemario := poemas.NewPoemario(predictor)

	// Operación
	poesia := poemario.GenerarPoesiaAPartirDe("fuego", "1", "3")

	// Validación
	palabrasDelVerso := poesia.Estrofas[0].Versos[0].Palabras
	if palabrasDelVerso[len(palabrasDelVerso)-1] == "a" {
		t.Logf("Se esperaba un verso sin conectores como palabra final pero se generó '%s'", poesia.Estrofas[0].Versos[0].Palabras)
		t.FailNow()
	}

}

func TestEsVacio(t *testing.T) {

	// Inicialización
	verso := new(poemas.Verso)

	// Operación
	esVacio := verso.EsVacio()

	// Validación
	if esVacio == false {
		t.Logf("Se esperaba true pero se obtuvo '%v'", esVacio)
		t.FailNow()
	}
}
