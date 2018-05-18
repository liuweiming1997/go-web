package filter

import (
	"fmt"
	"testing"
)

const (
	key = `httpS://tEStVIMIfsadf`
)

func TestReTitle(t *testing.T) {
	s := &VimiRegexp{
		BeginWith:   []string{"http://", "https://"},
		EndWith:     []string{"VIMI"},
		MustContain: []string{"test"},
	}

	re := s.GetRegexp()
	str := re.FindAllString(key, -1)
	fmt.Println(str)

	str2 := ReTest.FindAllString(key, -1)
	fmt.Println(str2)
}
