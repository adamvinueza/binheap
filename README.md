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
