package quicksort_test

import (
    "testing"
    "quicksort"
    "time"
    "math/rand"
)

// check if the slice comparison works correctly
func TestCompare(t *testing.T) {
    if !quicksort.Compare([]int{}, []int{}) {
        t.Error("Unexpected result, expected TRUE")
    }
    if quicksort.Compare([]int{1,2,3}, []int{1,2}) {
        t.Error("Unexpected result, expected FALSE")
    }

    if !quicksort.Compare([]int{1,2,3,4,5,6,7}, []int{1,2,3,4,5,6,7}) {
        t.Error("Unexpected result, expected TRUE")
    }
}

func TestQuickSort(t *testing.T)  {
    arr := []int{7,6,2,9,7,19,12,5}
    result := quicksort.QuickSort(arr)
    if (!quicksort.Compare(result,[]int{2,5,6,7,7,9,12,19})) {
        t.Error("result unexpected, ", result)
    }

    arr = []int{1,2,1,12821,12782,1289731,12,12987,1287,123987,5,12379}
    result = quicksort.QuickSort(arr)
    if !quicksort.Compare(result, []int{1,1,2,5,12,1287,12379,12782,12821,12987,123987,1289731}) {
        t.Error("result unexpected, ", result)
    }


    // more of a functional test - check if 10000 numbers will be sorted correctly
    rands := make([]int, 10000)
    s1 := rand.NewSource(time.Now().UnixNano())
    r1 := rand.New(s1)
    for i := 0; i < 10000; i++ {
        rands[i] = r1.Intn(1000000)
    }

    res := quicksort.QuickSort(rands)
    test := true
    for i, v := range res {
        // ignore 1st item
        if (i != 0) {
            if v < res[i-1] {
                test = false
                break
            }
        }
    }
    if !test {
        t.Error("Not all items were in order")
    }
}
