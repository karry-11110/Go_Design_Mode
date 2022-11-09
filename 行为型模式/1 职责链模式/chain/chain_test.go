package chain

import "testing"

func TestType(t *testing.T) {
	chain := &SensitiveWordFilterChain{}
	chain.AddFilter(&AdSensitiveWordFilter{})
	chain.AddFilter(&PoliticalWordFilter{})
	println(chain.Filter("test"))
}
