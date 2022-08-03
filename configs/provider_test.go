package configs

import "testing"

func TestGetDefaultProvider(t *testing.T) {
	p, err := GetDefaultProvider(ProviderAlicloud)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("provider: %v\n", p)
}
