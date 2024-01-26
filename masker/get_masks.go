package masker

func GetMasksLInks(s string, url string) string {
	arrWords := getArrWords(s)
	if (arrWords == nil) || (len(arrWords) == 0) {
		return s
	}
	var flag bool
	all := make([]byte, 0, 100)
	for _, word := range arrWords {
		strRune := []byte(word)
		if len(word) < len(url) {
			all = append(all, strRune...)
			continue
		}
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
	if len(s) == 0 {
		return nil
	}
	var result []string
	byteWord := make([]byte, 0, 100)
	for _, item := range s {
		byteWord = append(byteWord, byte(item))

		if (item == ' ') || (item == '\n') || (item == ',') || (item == '-') {
			result = append(result, string(byteWord))
			byteWord = nil
		}
	}
	result = append(result, string(byteWord))
	return result
}
