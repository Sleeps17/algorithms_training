package main

type pair struct {
	first  int64
	second int64
}

type pairSlice []pair

func (s pairSlice) Len() int {
	return len(s)
}

func (s pairSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s pairSlice) Less(i, j int) bool {
	if s[i].first == s[j].first {
		return s[i].second > s[j].second
	}

	return s[i].first < s[j].first
}

func main() {
	mainC()
}
