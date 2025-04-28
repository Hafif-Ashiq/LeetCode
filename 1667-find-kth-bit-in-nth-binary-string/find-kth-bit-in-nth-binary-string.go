func findKthBit(n, k int) byte {

	if k == 1 {
		return '0'
	}

	str := []byte{'0'}

	return letters(str, k, n)
}
func invert(arr []byte) []byte {
	newArr := make([]byte, len(arr))
	for i, v := range arr {
		
		if v == '1' {
			newArr[i] = '0'
		} else {
			newArr[i] = '1'
		}
	}
	return newArr
}

func letters(s []byte, k int, n int) byte {
	length := len(s)

	if length > k && n == 0 {
		
		return s[k-1]
	}
	second := invert(s)
	second = append(second, '1')

	slices.Reverse(second)

	new := slices.Concat(s, second)
	
	return letters(new, k, n-1)
}