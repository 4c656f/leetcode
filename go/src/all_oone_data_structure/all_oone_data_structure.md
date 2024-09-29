

----

Tags: #leetcode #design #hard #linkedList #hashMap

----

## Drawing:
[[all_oone_data_structure.excalidraw]]

----


# Intuition

we can use linked list as monotonic queue to achive O(1) complexity, cause of nature of inc/dec by one, we just need to swap self with neighb by right/left side

  

# Approach
use two datastructures: hashmap that will lead string to its node in linked queue, and queue as double linked list, keep track of head and tail of this list, thats will be our min and max,
on increment create new node to right/or if node with such count exist, simply add string to this node && update tail of list if its new max
same for dec

  

# Complexity

- Time complexity:

 $$O(1)$$

  

- Space complexity:

$$O(1)$$