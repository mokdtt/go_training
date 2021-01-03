package intset

import (
	"bytes"
	"fmt"
)

type MapIntSet struct {
	words map[int]bool
}

func (s *MapIntSet) Has(x int) bool {
	if s.words == nil {
		return false
	}
	return s.words[x]
}

func (s *MapIntSet) Add(x int) {
	if s.words == nil {
		s.words = map[int]bool{}
	}
	s.words[x] = true
}

func (s *MapIntSet) UnionWith(t *MapIntSet) {
	if s.words == nil {
		s.words = map[int]bool{}
	}
	if t.words == nil {
		return
	}
	for x, exists := range t.words {
		if exists {
			s.words[x] = true
		}
	}
}

func (s *MapIntSet) String() string {
	if s.words == nil {
		return ""
	}
	var buf bytes.Buffer
	buf.WriteByte('{')
	for word, exists := range s.words {
		if !exists {
			continue
		}
		if word != 0 {
			buf.WriteByte(' ')
		}
		fmt.Fprintf(&buf, "%d", word)
	}
	buf.WriteByte('}')
	return buf.String()
}
