package toolkitmodule

import "crypto/rand"

const randomString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

// ToolkitModule is a module that contains a set of functions that can be used in any other module
type ToolkitModule struct {
}

func (t *ToolkitModule) RandomString(n int) string {

	s, r := make([]rune, n), []rune(randomString)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
