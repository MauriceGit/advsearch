
// sortedSearch searches elements in sorted slices or user defined collections of any kind
// Because of the generic interfaces, sortedSearch can also be used to define possible insertion
// positions in a data structure or a position in a data structure based on other criteria
// based on the user's implementation of the Match() function.
package advsearch

import (
    "errors"
    "math"
)

const (
    EPS float64 = 0.0000001
)

// Implement the Searchable interface to be able to use BinarySearch
// on your data structure.
// The package requires that elements can be enumerated with an integer based index.
type Searchable interface {
    // True, if the value of e is smaller than the element at index i
    Smaller(e interface{}, i int) bool
    // Length of the data structure (element count)
    Len() int
    // Defines if we have a match based on the index i.
    // Most likely when e and value at i are equal.
    // You may also look at the elements before or after i to
    // determine a match (i.e. to find a possible insertion position).
    Match(e interface{}, i int) bool
}
// Implement the SearchableInterpolation interface to be able to use
// Interpolation Search and Quadratic Binary Search on your data structure.
type SearchableInterpolation interface {
    // We need all defined methods from Searchable
    Searchable
    // Method to get a value that can be used to interpolate the approximate
    // position in your data structure. Integers or other numbers can just
    // be casted to float64. This will not affect your search negatively.
    // This is needed for interpolation search and quadratic binary search.
    GetValue(i int) float64
    // Additionally, we don't know, what elements are in the data structure.
    // So we need the additional conversion function to work with e.
    ExtractValue(e interface{}) float64
}

func findLeftInterval(s SearchableInterpolation, e interface{}, stepSize, start int) int {
    thisStep := start-stepSize
    for s.ExtractValue(e) < s.GetValue(thisStep) && thisStep > 0 {
        thisStep -= stepSize
    }
    if thisStep < 0 {
        return 0
    }
    return thisStep
}

func findRightInterval(s SearchableInterpolation, e interface{}, stepSize, start int) int {
    thisStep := start+stepSize
    for thisStep < s.Len() && s.ExtractValue(e) > s.GetValue(thisStep) {
        thisStep += stepSize
    }
    if thisStep >= s.Len() {
        return s.Len()-1
    }
    return thisStep
}

func quadSearch(s SearchableInterpolation, e interface{}, l, r, count int) (int, int, error) {

    // Check against possible division by zero later.
    if s.Match(e, l) {
        return l, count, nil
    }
    if r < l || math.Abs(s.GetValue(l)-s.GetValue(r)) <= EPS {
        return -1, count, errors.New("No index could be found for the given element.")
    }

    percentage := (s.ExtractValue(e) - s.GetValue(l)) / (s.GetValue(r) - s.GetValue(l))

    // When the element we are looking for is outside the given range
    if percentage > 1.0 || percentage < 0.0 {
        return -1, count, errors.New("The element seems to be outside the given range.")
    }

    index := int(percentage * float64(r-l) + float64(l))

    if s.Match(e, index) {
        return index, count, nil
    }
    if s.Smaller(e, index) {
        return quadSearch(s, e, findLeftInterval(s, e, int(math.Sqrt(float64(index-1-l))), index-1), index-1, count+1)
    }
    return quadSearch(s, e, index+1, findRightInterval(s, e, int(math.Sqrt(float64(r-(index+1)))), index+1), count+1)
}

// QuadraticBinarySearch searches a match of the element e in the sorted data structure s.
// It will return the index/position in s based on your implementation of s.Match().
// Assuming that all methods of the interface Searchable and SearchableInterpolation run in O(1),
// QuadraticBinarySearch will have a time complexity of θ(sqrt(n)) worst case and θ(log(log(n))) average case.
// Pre-condition: s must be sorted!
func QuadraticBinarySearch(s SearchableInterpolation, e interface{}) (int, int, error) {
    return quadSearch(s, e, 0, s.Len()-1, 1)
}

func intSearch(s SearchableInterpolation, e interface{}, l, r, count int) (int, int, error) {

    // Check against possible division by zero later.
    if s.Match(e, l) {
        return l, count, nil
    }
    if r < l || math.Abs(s.GetValue(l)-s.GetValue(r)) <= EPS {
        return -1, count, errors.New("No index could be found for the given element.")
    }

    percentage := (s.ExtractValue(e) - s.GetValue(l)) / (s.GetValue(r) - s.GetValue(l))

    // When the element we are looking for is outside the given range
    if percentage > 1.0 || percentage < 0.0 {
        return -1, count, errors.New("The element seems to be outside the given range.")
    }

    index := int(percentage * float64(r-l) + float64(l))

    if s.Match(e, index) {
        return index, count, nil
    }
    if s.Smaller(e, index) {
        return intSearch(s, e, l, index-1, count+1)
    }
    return intSearch(s, e, index+1, r, count+1)

}

// InterpolationSearch searches a match of the element e in the sorted data structure s.
// It will return the index/position in s based on your implementation of s.Match().
// Assuming that all methods of the interface Searchable and SearchableInterpolation run in O(1),
// InterpolationSearch will have a time complexity of θ(n) worst case and θ(log(log(n))) average case.
// Pre-condition: s must be sorted!
func InterpolationSearch(s SearchableInterpolation, e interface{}) (int, int, error) {
    return intSearch(s, e, 0, s.Len()-1, 1)
}

func binSearch(s Searchable, e interface{}, l, r, count int) (int, int, error) {
    if r < l {
        return -1, count, errors.New("No index could be found for the given element in the data structure")
    }

    index := (r+l)/2

    if s.Match(e, index) {
        return index, count, nil
    }

    if s.Smaller(e, index) {
        return binSearch(s, e, l, index-1, count+1)
    }
    return binSearch(s, e, index+1, r, count+1)

}

// BinarySearch searches a match of the element e in the sorted data structure s.
// It will return the index/position in s based on your implementation of s.Match().
// Assuming that all methods of the interface Searchable run in O(1), BinarySearch
// will have a guaranteed time complexity of θ(log(n)) worst and average case.
// Pre-condition: s must be sorted!
func BinarySearch(s Searchable, e interface{}) (int, int, error) {
    return binSearch(s, e, 0, s.Len()-1, 1)
}
