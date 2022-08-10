package configs

import "testing"

func TestGetDefaultProvider(t *testing.T) {
	ps, err := GetDefaultProvider()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("providers: %v\n", ps)
}
