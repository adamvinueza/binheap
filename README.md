# Binary Heap

A treatment of binary heaps and their use in other data structures.

## Why are binary heaps useful?

Binary heaps are used in many other important data structures and algorithms.
For example, they are often used as implementations of _priority queues_,
because the items they store are sorted. And as priority queues are central to
many other important algorithms--such as Dijkstra's algorithm for discovering
the shortest path between two nodes in a graph--they can help us understand how
these more complex algorithms work.

## What is a binary heap?

A binary heap is a binary tree with two important properties. First, it is
_complete_: this means that every level of the tree is filled with the possible
exception of the lowest level, which is filled from left to right. Second, the
parents and children of the tree are ordered by a relation _R_ such that if _i_
is the parent of _j_, _R(i, j)_.

The second property is known as the _heap property_, and it's rather abstract,
so a couple of examples may help make it clear.

Suppose we're storing numbers in a binary heap and we want the numbers ordered
so that when we want to pull numbers out, we always get the largest one. Then if
_i_ and _j_ are stored in the heap and _i_ is the parent of _j_, _i_ cannot be
less than _j_.

Or suppose we're storing items from a grocery store so that the cheapest item is
always retrieved first. Then if _i_ and _j_ are items in the heap and _i_ is the
parent of _j_, then _i_ cannot be more expensive than _j_.

Or perhaps you have color samples and you want to always get the darkest sample
in the bunch. Then _i_ cannot to be lighter than _j_ if _i_ is _j_'s parent.

These examples show it doesn't matter _what_ you're ordering or what particular
order you care about--you can order those things using a binary heap.

## Applications of binary heaps

The examples above illustrated uses of a _priority queue_. A priority queue
isn't a queue, precisely--queues are structures wherein the order of retrieval
is identical to the order of insertion--although queues _are_ special cases of
priority queues (the ordering relation would be _earlier insertion_). Priority
queues are defined with an ordering relation, and their items can be retrieved
in that order.

As you may imagine, it's also possible to _sort_ using a binary heap. Create a
binary heap whose ordering relation is the one you would use to sort the items,
put the items into the heap, then extract them from the heap until it is empty;
the result is a list of items in the proper order.

## Efficiency of binary heaps

Binary heaps are efficient data structures. Priority queues can be created in
_O(n log n)_ time (on par with the most efficient sorting algorithms), and
extracting the first item can be done in _O(log n)_ time.
