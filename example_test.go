package advsearch_test

import (
    "fmt"
    "advsearch"
)

func ExampleBinarySearch() {
    s := []int{1, 2, 3, 4, 5, 6} // sorted ascending
    index, err := advsearch.BinarySearch(s, 4)
    fmt.Println(index)
    // Output: 3
}
func ExampleInterpolationSearch() {
    s := []int{1, 2, 3, 4, 5, 6} // sorted ascending
    index, err := advsearch.InterpolationSearch(s, 4)
    fmt.Println(index)
    // Output: 3
}
func ExampleQuadraticBinarySearch() {
    s := []int{1, 2, 3, 4, 5, 6} // sorted ascending
    index, err := advsearch.QuadraticBinarySearch(s, 4)
    fmt.Println(index)
    // Output: 3
}
