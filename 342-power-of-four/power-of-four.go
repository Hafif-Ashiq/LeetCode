func isPowerOfFour(n int) bool {
    return (n & (n - 1)) == 0 && (n-1) % 3 == 0 && n > 0
}