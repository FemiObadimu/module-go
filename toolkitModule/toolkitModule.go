package toolkitmodule

import "crypto/rand"

type ToolkitModule struct {
}

const randomString = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func (t *ToolkitModule) RandomString(n int) string {

	s, r := make([]rune, n), []rune(randomString)
	for i := range s {
		p, _ := rand.Prime(rand.Reader, len(r))
		x, y := p.Uint64(), uint64(len(r))
		s[i] = r[x%y]
	}

	return string(s)
}
