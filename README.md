# advsearch
A Go library for advanced searching in sorted data structures. The classic application would be
searching the index of an element in a sorted slice or verifying that an element exists in a slice.

By making this interface completely generic, neither the element type nor the data structure type
is fixed.

## Documentation:



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















### Data structure: []int, Element count: 1000, Averaged Test Cases: 1000000
| Algorithm | Tests until found| Average Time in ns |
| :--- | :--- | :--- |
| Binary search           | 8.52 | 385 |
| Interpolation search    | 2.96 | 352 |
| Quadratic binary search | 2.61 | 362 |

### Data structure: []int, Element count: 10000, Averaged Test Cases: 1000000
| Algorithm | Tests until found| Average Time in ns |
| :--- | :--- | :--- |
| Binary search           | 11.86 | 487 |
| Interpolation search    | 3.38  | 393 |
| Quadratic binary search | 2.93  | 415 |

### Data structure: []int, Element count: 100000, Averaged Test Cases: 1000000
| Algorithm | Tests until found| Average Time in ns |
| :--- | :--- | :--- |
| Binary search           | 15.18 | 618 |
| Interpolation search    | 3.64  | 420 |
| Quadratic binary search | 3.18  | 446 |

### Data structure: []int, Element count: 1000000, Averaged Test Cases: 1000000
| Algorithm | Tests until found| Average Time in ns |
| :--- | :--- | :--- |
| Binary search           | 18.45 | 835 |
| Interpolation search    | 4.08  | 499 |
| Quadratic binary search | 3.51  | 526 |

### Data structure: []int, Element count: 10000000, Averaged Test Cases: 1000000
| Algorithm | Tests until found| Average Time in ns |
| :--- | :--- | :--- |
| Binary search           | 21.82 | 1190 |
| Interpolation search    | 4.25  | 578  |
| Quadratic binary search | 3.69  | 617  |

### Data structure: []int, Element count: 100000000, Averaged Test Cases: 1000000
| Algorithm | Tests until found| Average Time in ns |
| :--- | :--- | :--- |
| Binary search           | 25.14 | 1624 |
| Interpolation search    | 4.52  | 636  |
| Quadratic binary search | 3.88  | 698  |
