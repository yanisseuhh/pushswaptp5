package main

import (
	"testing"
)

func TestParseArgs(t *testing.T) {
	cas := []struct {
		entree      string
		doitEchouer bool
	}{
		{"2 1 3", false},
		{"-5 0 7", false},
		{"0 one 2 3", true},
		{"1 2 2", true},
		{"+ 3", true},
		{"2147483648", true},
	}

	for _, c := range cas {
		_, err := parseArgs([]string{c.entree})
		if c.doitEchouer && err == nil {
			t.Errorf("%q : une erreur était attendue", c.entree)
		}
		if !c.doitEchouer && err != nil {
			t.Errorf("%q : erreur inattendue : %v", c.entree, err)
		}
	}
}

func TestOperations(t *testing.T) {

	a := []int{2, 1, 3}
	a, _ = appliquer("sa", a, []int{})
	if a[0] != 1 || a[1] != 2 {
		t.Errorf("sa échoué : %v", a)
	}

	a, b := appliquer("pb", []int{1, 2}, []int{})
	if len(b) != 1 || b[0] != 1 || len(a) != 1 {
		t.Errorf("pb échoué")
	}

	a, _ = appliquer("ra", []int{1, 2, 3}, []int{})
	if a[0] != 2 || a[2] != 1 {
		t.Errorf("ra échoué")
	}

	a, b = appliquer("pa", []int{1, 2}, []int{})
	if len(b) != 0 || len(a) != 2 {
		t.Errorf("pa sur pile vide échoué")
	}
}

func TestTriExempleSujet(t *testing.T) {
	depart := []int{2, 1, 3, 6, 5, 8}
	ops := forceBrute(depart)
	if len(ops) >= 9 {
		t.Errorf("Attendu moins de 9 opérations, obtenu : %d", len(ops))
	}
}
