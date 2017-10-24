package gcd

// Use Euclid's algorithm to calculate the GCD of 2 numbers
func GreatestCommonDenominator(a, b int) int {
    for b != 0 {
        t := b
        a, b = t, a % b
    }
    return a
}
