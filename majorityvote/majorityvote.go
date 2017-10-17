package main

import "fmt"

// perform the Boyer-Moore Majority Vote algorithm
func majorityVote(arr []string) string {
    candidate := ""
    count := 0
    // perform the initial algorithm
    for _, v := range arr {
        if count == 0 {
            candidate = v
        }
        if candidate == v {
            count++
        } else {
            count--
        }
    }
    // perform a check to see if they are the majority
    totalCount := 0
    for _, v := range arr {
        if v == candidate {
            totalCount++
        }
    }
    if totalCount > len(arr)/2 {
        return candidate
    } else {
        return ""
    }
}

func main() {
    arr := []string{"a", "b", "c", "a", "a", "b", "d", "a", "a"}
    fmt.Printf("The majority vote of %s is '%s'.\n", arr, majorityVote(arr))

    arr2 := []string{"a", "b", "c", "a", "a", "b", "d"}
    fmt.Printf("The majority vote of %s is '%s'.\n ", arr2, majorityVote(arr2))

}
