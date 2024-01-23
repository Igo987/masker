package main

func GetMasksLInks(s string, url string) string {
	arrWords := getArrWords(s)
	flag := false
	all := make([]byte, 0, 100)
	for _, word := range arrWords {
		strRune := []byte(word)
		for i := 0; i < len(url); i++ {
			if url[i] == strRune[i] {
				flag = true
			} else {
				all = append(all, strRune...)
				flag = false
				break
			}
		}
		if flag {
			for i := len(url); i < len(strRune); i++ {
				strRune[i] = '*'
			}
			all = append(all, strRune...)
		}
	}
	return string(all)
}

func getArrWords(s string) []string {
	var result []string
	byteWord := make([]rune, 0, 10)
	for _, item := range s {
		byteWord = append(byteWord, item)
		if (item == ' ') || (item == '\n') || (item == ',') || (item == '-') {
			result = append(result, string(byteWord))
			byteWord = nil
		}
	}
	return result
}
