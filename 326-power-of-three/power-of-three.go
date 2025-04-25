func isPowerOfThree(n int) bool {
    
    if n % 3 == 0 && n != 0 {
        if n / 3 == 1{
            return true
        }
        return isPowerOfThree(n/3)
    } else if n == 1{
        return true
    } else{
        return false
    }
}