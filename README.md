# advsearch
A Go library for advanced searching in sorted data structures. The classic application would be
searching the index of an element or verifying that an element exists in a sorted slice.

By making this interface completely generic, neither the element type nor the data structure type
is predetermined. This way it is very easy to use i.e. InterpolationSearch on a slice of complex structs or
any other kind of sorted collection with nearly any kind of data!

The choice of interface definition and functionality are oriented and influenced by the [*sort* library](https://golang.org/pkg/sort/).

## Documentation:

Three functions are defined for searching in sorted data structures:

### _Binary Search_:

Binary search guarantees a runtime within θ(log(n)) (worst and average case) and is the standard method used for finding
an element in a sorted slice (Binary search for a few standard data types like int/float32 is already implemented in the *sort* library!).

If you are in an academic field and have to guarantee logarithmic runtime, you should use Binary search!

### _Interpolation Search_:

Interpolation search has an improved average case runtime of θ(log(log(n))) but a worst case runtime of θ(n).
If elements are uniformly distributed within your data structure, Interpolation search will yield much better
results than Binary search. For highly irregular data, I recommend to try it in your special test case and
possibly move to Quadratic Binary search, if the worst case kicks in too often.

### _Quadratic Binary Search_:

Quadratic Binary search improves the worst case runtime of Interpolation search to θ(sqrt(n)) while keeping the average
runtime at θ(log(log(n))). If your data is very uniformly distributed, it will be a little slower than Interpolation searching
because of a slight internal overhead (see Performance measurements). In all other cases and generally more irregular data distribution,
Quadratic Binary search will perform better than Interpolation search both in time the function takes and how many elements
need to be compared to find the correct one.

## Usage:

To be able to use the BinarySearch, your data structure has to implement the *Searchable* interface (defined in advsearch.go).

The *Searchable* interface defines the following methods:

```go
Len() int
```
>Returns the length/element count of the data structure.
```go
Smaller(e interface{}, i int) bool
```
>e can be savely casted to the element type you use in your data structure!
>Smaller() should return True, if the value of e is smaller than the element at index i in your data structure.

```go
Match(e interface{}, i int) bool
```
>e can be savely casted to the element type you use in your data structure!
>Match defines if we have a match based on the index i.
>When searching the index of an element e, Match() should return True, when e and value at i are equal.
>You may also look at the elements before or after i to determine a match (i.e. to find a possible insertion position).

To be able to use InterpolationSearch and QuadraticBinarySearch, your data structure additionally has to implement the *SearchableInterpolation* interface (defined in advsearch.go).

The *SearchableInterpolation* interface defines the following methods:

```go
GetValue(i int) float64
```
>Method to get a value at index i, that can be used to interpolate the approximate position in your data structure. Integers or other numbers can just
>be casted to float64. This will not affect your search negatively.

```go
ExtractValue(e interface{}) float64
```
>e can be savely casted to the element type you use in your data structure!
>ExtractValue extracts/converts a value from the given element e. Integers or other numbers can just
>be casted to float64. This will not affect your search negatively.

A fully functional example of finding the index of an element in a sorted slice:

```go
package main
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

func main() {
    s := IntSlice{1, 2, 3, 4, 5, 6} // sorted ascending
    index, _ := advsearch.QuadraticBinarySearch(&s, 4)
    fmt.Println(index)
    // Output: 3
}
```

## Performance measurements:

Data structure: []int, Averaged Test Cases: 1000000

### Tests until a random element is found in a slice of *n* elements:
| Algorithm               | n = 1000 | n = 10000 | n = 100000 | n = 1000000 | n = 10000000 | n = 100000000 |
| :---                    | :---     | :---      | :---       | :---        | :---         | :---          |
| Binary search           | 8.52     | 11.86     | 15.18      | 18.45       | 21.82        | 25.14         |
| Interpolation search    | 2.96     | 3.38      | 3.64       | 4.08        | 4.25         | 4.52          |
| Quadratic binary search | 2.61     | 2.93      | 3.18       | 3.51        | 3.69         | 3.88          |

### Time (in nanoseconds) until a random element is found in a slice of *n* elements:
| Algorithm               | n = 1000 | n = 10000 | n = 100000 | n = 1000000 | n = 10000000 | n = 100000000 |
| :---                    | :---     | :---      | :---       | :---        | :---         | :---          |
| Binary search           | 385      | 487       | 618        | 835         | 1190         | 1624          |
| Interpolation search    | 352      | 393       | 420        | 499         | 578          | 636           |
| Quadratic binary search | 362      | 415       | 446        | 526         | 617          | 698           |
