package factory

import "testing"

func TestBinario_Convertir(t *testing.T) {
	Binario := NewBinario()

	testCases := []struct {
		name    string
		texto   string
		wantErr bool
	}{
		{name: "Test Success ",
			texto:   "Alexander",
			wantErr: false,
		},
	}

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			_, err := Binario.Convertir(test.texto)
			if err != nil && !test.wantErr {
				t.Fatalf("Convertir function Error %v received", err)
			}
		})
	}
}
