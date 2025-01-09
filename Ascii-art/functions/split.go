package ascii

func Split(str string) []string {
	slice := []string{}
	newStr := ""

	for i := 0; i < len(str); i++ {
		if i < len(str)-1 && str[i] == '\\' && str[i+1] == 'n' {
			if newStr != "" {
				slice = append(slice, newStr)
				newStr = ""
			}
			slice = append(slice, "")
			i += 1
		} else {
			newStr += string(str[i])
		}
	}

	if newStr != "" {
		slice = append(slice, newStr)
	}
	return slice
}
