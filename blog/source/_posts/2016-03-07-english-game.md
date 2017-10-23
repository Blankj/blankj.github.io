---
title: English Game
date: 2016-03-07 20:16:46
categories:
  - OJ
tags:
  - 树
  - DP
---

题目描述

This English game is a simple English words connection game.

The rules are as follows: there are N English words in a dictionary, and every word has its own weight v. There is a weight if the corresponding word is used. Now there is a target string X. You have to pick some words in the dictionary, and then connect them to form X. At the same time, the sum weight of the words you picked must be the biggest.
<!--more-->

输入要求

There are several test cases. For each test, N (1<=n<=1000) and X (the length of x is not bigger than 10000) are given at first. Then N rows follow. Each row contains a word wi (the length is not bigger than 30) and the weight of it. Every word is composed of lowercases. No two words in the dictionary are the same.

输出要求

For each test case, output the biggest sum weight, if you could not form the string X, output -1.

假如输入
<pre>1 aaaa
a 2
3 aaa
a 2
aa 5
aaa 6
4 abc
a 1
bc 2
ab 4
c 1
3 abcd
ab 10
bc 20
cd 30
3 abcd
cd 100
abc 1000
bcd 10000
</pre>
应当输出
<pre>8
7
5
40
-1
</pre>


Java：

``` java
import java.io.BufferedInputStream;
import java.util.Scanner;

public class Main {
    // 字典树dp
    static class Node {
        int w;
        boolean end;
        Node[] next = new Node[26];
    }

    static int[] dp;
    static String str;
    static Node root;
    static int len;

    public static void main(String[] args) {
        Scanner scanner = new Scanner(new BufferedInputStream(System.in));
        while (scanner.hasNext()) {
            root = new Node();
            int n = scanner.nextInt();
            str = " ";
            str += scanner.next();
            len = str.length();
            dp = new int[len];
            for (int i = 0; i < n; ++i) {
                String word = scanner.next();
                int w = scanner.nextInt();
                insertNode(word, w);
            }
            solve(0);
            for (int i = 1; i < len - 1; ++i) {
                if (dp[i] != 0)
                    solve(i);
            }
            System.out.println(dp[len - 1] == 0 ? -1 : dp[len - 1]);
        }

    }

    static void insertNode(String word, int w) {
        int wlen = word.length();
        Node tmp = root;
        for (int i = 0; i < wlen; ++i) {
            int x = word.charAt(i) - 'a';
            if (tmp.next[x] == null) {
                tmp.next[x] = new Node();
            }
            tmp = tmp.next[x];
        }
        tmp.end = true;
        tmp.w = w;
    }

    static void solve(int i) {
        Node cur = root;
        for (int j = i + 1; j < len; ++j) {
            int x = str.charAt(j) - 'a';
            if (cur.next[x] != null) {
                if (cur.next[x].end) {
                    dp[j] = Math.max(dp[j], dp[i] + cur.next[x].w);
                }
                cur = cur.next[x];
            } else {
                break;
            }
        }
    }
}
```