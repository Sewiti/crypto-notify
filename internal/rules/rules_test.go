package rules

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	testFile    = "testdata/test.json"
	badTestFile = "testdata/bad.json"
)

func TestReadWrite(t *testing.T) {
	rules, err := Read(testFile)
	if err != nil {
		t.Fatal(err)
	}

	tmpFile, err := ioutil.TempFile(os.TempDir(), "cryptonotify_rule_test_*.json")
	if err != nil {
		t.Fatal(err)
	}

	defer os.Remove(tmpFile.Name())

	err = Write(tmpFile.Name(), rules)
	if err != nil {
		t.Fatal(err)
	}

	rules2, err := Read(tmpFile.Name())
	if err != nil {
		t.Fatal(err)
	}

	for i := range rules {
		if rules[i] != rules2[i] {
			t.Fatal("Read & written files don't match")
		}
	}
}

func TestReadFailure(t *testing.T) {
	_, err := Read(badTestFile)
	if err == nil {
		t.Fatal("Expected error reading bad file")
	}
}

func TestTableCheck(t *testing.T) {
	tests := []struct {
		input    float64
		operator string
		target   float64
		expected bool
	}{
		// Greater than
		{500, "gt", 400, true},
		{500, "Gt", 500, false},
		{500, "gT", 600, false},
		{500.01, "GT", 500, true},

		// Greater than or equals
		{500, "ge", 400, true},
		{500, "Ge", 500, true},
		{500, "gE", 500.01, false},
		{499.69, "GE", 499.70, false},

		// Less than
		{400, "lt", 500, true},
		{500, "LT", 500, false},
		{600, "Lt", 500, false},
		{499.99, "lT", 500, true},

		// Less than or equals
		{400, "le", 500, true},
		{500.1, "lE", 500.1, true},
		{600, "Le", 500, false},
		{499.99, "LE", 500, true},

		// Equals
		{499.99, "eq", 500, false},
		{500.1, "eQ", 500.1, true},
		{600, "Eq", 500, false},
		{549.51, "EQ", 549.51, true},

		// Not equals
		{499.99, "ne", 500, true},
		{245.08, "nE", 245.08, false},
		{600, "Ne", 700, true},
		{349.13, "NE", 349.13, false},
	}

	for _, test := range tests {
		r := Rule{
			Price: test.target,
			Op:    test.operator,
		}

		trig, err := r.Check(test.input)
		if err != nil {
			t.Fatal(err)
		}

		if trig != test.expected {
			t.Fatalf("%f %s %f returned %t, expected %t", test.input, strings.ToLower(r.Op), r.Price, trig, test.expected)
		}
	}
}
