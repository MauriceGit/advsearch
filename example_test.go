package advsearch_test

import (
    "fmt"
    "github.com/MauriceGit/advsearch"
)

type IntSlice []int

// Len is part of Searchable.
func (s *IntSlice) Len() int {
    return len(s)
}
// Smaller is part of the Searchable interface
func (s *IntSlice) Smaller(e interface{}, i int) bool {
    return e.(int) < (*s)[i]
}
// Match is part of the Searchable interface
func (s *IntSlice) Match(e interface{}, i int) bool {
    return e.(int) == (*s)[i]
}
// GetValue is part of the SearchableInterpolation interface
func (s *IntSlice) GetValue(i int) float64 {
    return float64((*s)[i])
}
// ExtractValue is part of the SearchableInterpolation interface
func (s *IntSlice) ExtractValue(e interface{}) float64 {
    return float64(e.(int))
}

func ExampleBinarySearch() {
    s := IntSlice{1, 2, 3, 4, 5, 6} // sorted ascending
    index, _ := advsearch.BinarySearch(&s, 4)
    fmt.Println(index)
    // Output: 3
}
func ExampleInterpolationSearch() {
    // Output: 3
    s := IntSlice{1, 2, 3, 4, 5, 6} // sorted ascending
    index, _ := advsearch.InterpolationSearch(&s, 4)
    fmt.Println(index)
    // Output: 3
}
func ExampleQuadraticBinarySearch() {
    s := IntSlice{1, 2, 3, 4, 5, 6} // sorted ascending
    index, _ := advsearch.QuadraticBinarySearch(&s, 4)
    fmt.Println(index)
    // Output: 3
}
