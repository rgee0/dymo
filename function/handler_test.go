package function

import (
	"testing"
)

func Test_reverse(t *testing.T) {

	var TestCases = []struct {
		Name            string
		InputDomain     string
		ExpectedOuptput string
	}{
		{
			Name:            "Not a domain",
			InputDomain:     "Mary had a little lamb",
			ExpectedOuptput: "Mary had a little lamb",
		},
		{
			Name:            "Valid domain with http scheme",
			InputDomain:     "http://blog.openfaas.com",
			ExpectedOuptput: "com.openfaas.blog",
		},
		{
			Name:            "Valid domain with https scheme",
			InputDomain:     "https://blog.openfaas.com",
			ExpectedOuptput: "com.openfaas.blog",
		},
		{
			Name:            "Valid domain without scheme",
			InputDomain:     "rgee0.o6s.io",
			ExpectedOuptput: "io.o6s.rgee0",
		},
		{
			Name:            "Longer valid domain without scheme",
			InputDomain:     "mary.had.a.little.lamb.com",
			ExpectedOuptput: "com.lamb.little.a.had.mary",
		},
		{
			Name:            "Valid domain with http scheme with trailing path",
			InputDomain:     "http://blog.openfaas.com/mary/had/a/little/lamb",
			ExpectedOuptput: "com.openfaas.blog",
		},
		{
			Name:            "Valid domain with https scheme with trailing path",
			InputDomain:     "https://blog.openfaas.com/mary/had/a/little/lamb",
			ExpectedOuptput: "com.openfaas.blog",
		},
		{
			Name:            "Valid domain without scheme with trailing path",
			InputDomain:     "rgee0.o6s.io/mary/had/a/little/lamb",
			ExpectedOuptput: "io.o6s.rgee0",
		},
	}

	for _, test := range TestCases {
		res := reverse(test.InputDomain)

		if res != test.ExpectedOuptput {
			t.Errorf("Testcase %s failed. want - %s, got - %s", test.Name, test.ExpectedOuptput, res)
		}
	}
}
