package service

type Service struct {
}

func New() *Service {
	return &Service{}
}

func (s *Service) LongestSubstring(str string) string {

	// Создаем мапу для отслеживания индексов символов
	indexMap := make(map[byte]int)

	// Переменные для отслеживания текущей подстроки
	start := 0
	longestSubstring := ""

	// Перебираем символы строки
	for i := 0; i < len(str); i++ {
		// Если символ уже встречался и его индекс находится в текущей подстроке
		if index, ok := indexMap[str[i]]; ok && index >= start {
			// Обновляем начальный индекс текущей подстроки
			start = index + 1
		}

		// Обновляем индекс символа в мапе
		indexMap[str[i]] = i

		// Обновляем текущую подстроку, если она является самой длинной
		substring := str[start : i+1]
		if len(substring) > len(longestSubstring) {
			longestSubstring = substring
		}
	}

	return longestSubstring
}
