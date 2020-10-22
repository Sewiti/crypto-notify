package main

import (
	"context"
	"io/ioutil"
	"os"
	"testing"

	"github.com/Sewiti/crypto-notify/internal/rules"
)

type testFiles struct {
	input    string
	expected string
}

func TestIntegration(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test")
	}

	tests := []string{
		"testdata/test_1.json",
		"testdata/test_2.json",
	}

	for _, file := range tests {
		performTest(t, file)
	}
}

// prepareFile resets rules triggers and writes it to a temporary file
func prepareFile(t *testing.T, r rules.Rules) (temp *os.File) {
	re := make([]rules.Rule, len(r))
	copy(re, r)

	for i := range re {
		re[i].Triggered = false
	}

	temp, err := ioutil.TempFile(os.TempDir(), "cryptonotify_integration_test_*.json")
	if err != nil {
		t.Fatal(err)
	}

	err = rules.Write(temp.Name(), re)
	if err != nil {
		t.Fatal(err)
	}

	return temp
}

func performTest(t *testing.T, ruleFile string) {
	expected, err := rules.Read(ruleFile)
	if err != nil {
		t.Fatal(err)
	}

	temp := prepareFile(t, expected).Name()
	defer os.Remove(temp)

	ctx, cancel := context.WithTimeout(context.Background(), reqTimeout)
	defer cancel()

	err = exec(ctx, temp)
	if err != nil {
		t.Fatal(err)
	}

	r, err := rules.Read(temp)
	if err != nil {
		t.Fatal(err)
	}

	// Compare result to expected
	for i, v := range r {
		if v != expected[i] {
			t.Fatalf(
				"rule: %s returned %t, expected %t\n",
				v.String(),
				v.Triggered,
				expected[i].Triggered,
			)
		}
	}
}
