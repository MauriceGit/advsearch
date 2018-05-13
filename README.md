# advsearch
A Go library for advanced searching in sorted data structures. The classic application would be
searching the index of an element or verifying that an element exists in a sorted slice.

By making this interface completely generic, neither the element type nor the data structure type
is predetermined. The choice of interface definition and functionality are oriented and influenced by the *sort* library.

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
