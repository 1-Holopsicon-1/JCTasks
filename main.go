package main

func main() {

}

type StringIntMap struct {
	mapa map[string]int
}

func (s *StringIntMap) NewStringIntMap() {
	s.mapa = make(map[string]int)
}

func (s *StringIntMap) Add(key string, value int) {
	s.mapa[key] = value
}

func (s *StringIntMap) Remove(key string) {
	delete(s.mapa, key)
}

func (s *StringIntMap) Copy() (copied map[string]int) {
	for i, v := range s.mapa {
		copied[i] = v
	}
	return
}

func (s *StringIntMap) Exists(key string) bool {
	_, ok := s.mapa[key]
	return ok
}

func (s *StringIntMap) Get(key string) (int, bool) {
	v, ok := s.mapa[key]
	return v, ok
}
