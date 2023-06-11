package service

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) LongestSubstring(str string) string {

	index := make(map[byte]int)

	start := 0
	longestSubstring := ""

	for i := 0; i < len(str); i++ {

		if index, ok := index[str[i]]; ok && index >= start {
			start = index + 1
		}
		index[str[i]] = i
		substring := str[start : i+1]
		if len(substring) > len(longestSubstring) {
			longestSubstring = substring
		}
	}

	return longestSubstring
}
