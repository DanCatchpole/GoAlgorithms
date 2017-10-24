package gcd_test

import (
    "testing"
    "gcd"
)

func TestGreatestCommonDenominator(t *testing.T)  {
    a := gcd.GreatestCommonDenominator(182, 763)
    if (a != 7) {
        t.Error("Got ", a, "Expected 7.")
    }
    a = gcd.GreatestCommonDenominator(1287128,835832)
    if (a != 8) {
        t.Error("Got ", a, "Expected 8.")
    }
}
