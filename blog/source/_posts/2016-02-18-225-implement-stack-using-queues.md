---
title: 225. Implement Stack using Queues
date: 2016-02-18 16:20:44
categories:
  - LeetCode
  - Java
tags:
  - Stack
  - Design
---

Total Accepted: 31239
Total Submissions: 102745
Difficulty: Easy

Implement the following operations of a stack using queues.

*   push(x) -- Push element x onto stack.</p>
*   pop() -- Removes the element on top of the stack.

*   top() -- Get the top element.

*   empty() -- Return whether the stack is empty.

**Notes:**

*   You must use only standard operations of a queue -- which means only `push to back`, `peek/pop from front`, `size`, and `is empty` operations are valid.

*   Depending on your language, queue may not be supported natively. You may simulate a queue by using a list or deque (double-ended queue), as long as you use only standard operations of a queue.

*   You may assume that all operations are valid (for example, no pop or top operations will be called on an empty stack).

<!-- more -->

Java:

``` java
class MyStack {
    private Deque<Integer> queue = new ArrayDeque<>();

    // Push element x onto stack.
    public void push(int x) {
        queue.add(x);
    }

    // Removes the element on top of the stack.
    public void pop() {
        queue.removeLast();
    }

    // Get the top element.
    public int top() {
        return queue.getLast();
    }

    // Return whether the stack is empty.
    public boolean empty() {
        return queue.size()==0;
    }
}
```