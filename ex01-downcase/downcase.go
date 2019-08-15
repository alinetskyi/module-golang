package downcase

func Downcase(str string) (string, error) {
	tmp := ""
	for i := 0; i < len(str); i++ {
		char := str[i]
		if char < 91 && char > 64 {
			char = char + 32
		}
		tmp += string(char)
	}
	return tmp, nil
}
