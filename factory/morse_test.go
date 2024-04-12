package factory

import "testing"

func TestMorse_Convertir(t *testing.T) {
	Morse := NewMorse()

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
			_, err := Morse.Convertir(test.texto)
			if err != nil && !test.wantErr {
				t.Fatalf("Convertir function Error %v received", err)
			}
		})
	}
}
