

----

Tags: #leetcode #dp #sorting #hard

----

## Drawing:
[[minimum_total_distance_traveled.excalidraw]]

----


# Intuition

main idea is to how transform this problem into dp subproblems, obvious soluition is to write a recursion that will try all possible factories for all robots, that will lead us to $$O(2^(n+k))$$ time complexity 
but how we can transform this problem, so we could cache each recursion step? transform input factories into sorted array of factories where each factory is unique factory that have capacity of 1, so when we visit factory at index i it can be marked as used
so we can make decision tree for each robot, where we can repair it at i'th factory, or do not repair it, and cache min result of this two results in dp table where key to cache value is index of robot and index of current factory

# Approach

<!-- Describe your approach to solving the problem. -->

  

# Complexity

- Time complexity:

 $$O(n*k)$$

  

- Space complexity:

$$O(n*k)$$