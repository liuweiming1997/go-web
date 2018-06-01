package filter

import (
	"fmt"
	"testing"
)

const (
	key  = `httpS://tEStVIMIfsadf`
	key1 = `<a target="_blank" class="archive-title" href="http://blog.jobbole.com/113230/" title="客户想要的 vs 客户实际预算：漫画解读软件开发模式 &#8203;&#8203;&#8203;&#8203;">客户想要的 vs 客户实际预算：漫画解读软件开发模式 &#8203;&#8203;&#8203;&#8203;</a>`
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

func TestBoLeZaiXian(t *testing.T) {
	fmt.Println("TestBoLeZaiXian")
	s := &VimiRegexp{
		BeginWith:   []string{"http://"},
		EndWith:     []string{`"`},
		MustContain: []string{"blog.jobbole.com"},
	}
	re := s.GetRegexp()

	str := re.FindAllString(key1, -1)
	fmt.Println(str)
}
