---
title: 232. Implement Queue using Stacks
date: 2016-02-13 17:11:01
categories:
  - LeetCode
  - Java
tags:
  - Stack
  - Design
---

Total Accepted: **33142**
Total Submissions: **97820**
Difficulty: **Easy**

Implement the following operations of a queue using stacks.</p>

*   push(x) -- Push element x to the back of queue.</p>
*   pop() -- Removes the element from in front of queue.

*   peek() -- Get the front element.

*   empty() -- Return whether the queue is empty.

**Notes:**

*   You must use _only_ standard operations of a stack -- which means only `push to top`, `peek/pop from top`, `size`, and `is empty` operations are valid.

*   Depending on your language, stack may not be supported natively. You
 may simulate a stack by using a list or deque (double-ended queue), as
 long as you use only standard operations of a stack.

*   You may assume that all operations are valid (for example, no pop or peek operations will be called on an empty queue).

<!-- more -->

两个栈来实现队列的效果

Java:

``` java
class MyQueue {
    Stack<Integer> input = new Stack();
    Stack<Integer> output = new Stack();

    // Push element x to the back of queue.
    public void push(int x) {
        input.push(x);
    }

    // Removes the element from in front of queue.
    public void pop() {
        peek();
        output.pop();
    }

    // Get the front element.
    public int peek() {
        if (output.empty())
            while (!input.empty())
                output.push(input.pop());
        return output.peek();
    }

    // Return whether the queue is empty.
    public boolean empty() {
         return input.empty() && output.empty();
    }
}
```