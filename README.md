# advsearch
A Go library for advanced searching in sorted data structures




## Performance measurements

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
