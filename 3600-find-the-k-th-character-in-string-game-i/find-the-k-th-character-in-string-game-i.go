func kthCharacter(k int) byte {
	str := []byte{'a'}

	return letters(str, k)
}

func letters(s []byte, k int) byte {
	length := len(s)

	if length > k {
		return s[k-1]
	}

	for _, v := range s {
		if v == 'z' {
			s = append(s, 'a')
		} else {
			s = append(s, v+1)
		}
	}

	return letters(s, k)
}