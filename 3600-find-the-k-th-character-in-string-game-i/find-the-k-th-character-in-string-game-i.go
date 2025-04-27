func kthCharacter(k int) byte {
	str := "a"

	for len(str) <= k {
		newStr := str
		for _, v := range str {
			if v == 'z' {
				newStr += "a"
			} else {
				newStr += string(v + 1)
			}
		}
		str = newStr

	}
	return str[k-1]
}