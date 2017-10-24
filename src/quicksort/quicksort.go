package quicksort

// Sorts a slice of integers using the QuickSort algorithm
func QuickSort(arr []int) []int  {
    // if there is a single element then there is no need to
    if (len(arr) == 1 || len(arr) == 0) {
        return arr
    }

    // choose a pivot
    pivot := len(arr) / 2
    pivotItem := arr[pivot]

    // create slices to store the left+right of the array
    less := make([]int, 0)
    more := make([]int, 0)

    /* move the item to the less or more slice if they are lessthan
       or greaterorequal to the pivot respectively */
    for i, v := range arr {
        if i != pivot {
            if v < pivotItem {
                less = append(less, v)
            } else {
                more = append(more, v)
            }
        }
    }

    // quicksort the items less than and more than separately
    less = QuickSort(less)
    more = QuickSort(more)

    // combine the slices
    less = append(less, pivotItem)
    less = append(less, more...)

    // return the sorted slice
    return less
}

// Compares two slices to check if they are equal
func Compare(a, b []int) bool {
    if len(a) != len(b) {
        return false
    }
    for i, v := range a {
        if v != b[i] {
            return false
        }
    }
    return true
}
