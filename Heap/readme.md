# Heap

This project implement a simple Max heap, Min heap Min-Max heap use golang.

## Output
```
MaxHeap :
27      7       80      5       67      18      62      24      58      25
swap    5       58
swap    7       67
swap    7       25
swap    27      80
swap    27      62
80      67      62      58      25      18      27      24      5       7
MaxHeap Add and Remove :
40      23      10      15      8
40      23      10      15      8
Add : 60
swap    10      60
swap    40      60
60      23      40      15      8       10
Add : 20
60      23      40      15      8       10      20
Remove : 60
swap    60      20
swap    20      40
40      23      20      15      8       10
Remove : 23
swap    23      10
swap    10      15
40      15      20      10      8
MinHeap :
20      30      10      50      60      40      45      5       15      25
swap    60      25
swap    50      5
swap    30      5
swap    30      15
swap    20      5
swap    20      15
5       15      10      20      25      40      45      50      30      60
Min-MaxHeap :
20      8       28      10      4       5       40      55
swap    40      28
swap    55      8
swap    8       10
swap    4       20
4       55      40      8       20      5       28      10
Min-MaxHeap Add and Remove :
Add : 2
swap    2       8
swap    2       4
2       55      40      4       20      5       28      10      8
Remove : 40
swap    40      8
swap    28      8
2       55      28      4       20      5       8       10
```