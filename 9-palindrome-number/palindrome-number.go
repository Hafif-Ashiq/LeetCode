func isPalindrome(x int) bool {
    if x==0{
		return true
	}
	if x < 0 ||  x % 10 == 0{
		return false
	}
	if x < 10 {
		return true
	}

	temp := x
	reverse := 0
	//first := 1
	for temp >= 1 {
		rem := temp % 10
		temp = temp / 10
		// fmt.Println(rem)
		// fmt.Println(temp)
		// fmt.Println("--------")

		reverse = rem + (reverse * 10)
		//first *= 10

	}

	// fmt.Println(reverse)

	return reverse == x

}