package toolkitmodule

import "testing"

func TestTools_RandomString(t *testing.T) {
	var randomStrTool ToolkitModule

	s := randomStrTool.RandomString(5)

	if len(s) != 5 {
		t.Errorf("RandomString() = %v; want 5", len(s))
	}

}
