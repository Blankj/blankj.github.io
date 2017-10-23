---
title: Java对结构体的排序
date: 2015-04-01 19:28:50
categories:
  - OJ
tags:
  - 排序
---

### 题目描述

输入若干个学生的信息（学号、姓名、成绩），输入学号为0时结束，用单向链表组织这些学生的信息后，再按成绩由低到高顺序输出。

<!-- more -->
### 输入要求

每行输入若干个学生的信息（学号、姓名、成绩），学号和成绩为整数，姓名为长度不超过10的字符串。

输入学号为0时结束。

### 输出要求

按成绩由低到高输出所有学生的信息(如果成绩相同，按输入的次序输出)。

### 假如输入
<pre>1 zhang 78
2 wang 80
3 Li 75
4 zhao 85
0</pre>

### 应当输出
<pre>3 Li 75
1 zhang 78
2 wang 80
4 zhao 85</pre>


``` java
import java.util.Arrays;
import java.util.Comparator;
import java.util.Scanner;

class Stu {
	int index, id, score;
	String name;
}

class mycmp implements Comparator<Stu> {
	public int compare(Stu s1, Stu s2) {
		if (s1.score != s2.score)
			return s1.score - s2.score;
		else
			return s1.index - s2.index;
	}
}

public class Main {
	public static void main(String[] args) {
		Scanner scan = new Scanner(System.in);
		Stu[] s = new Stu[1000];
		int temp, i;
		for (i = 0; i < 1000; i++) {
			temp = scan.nextInt();
			if (temp == 0)
				break;
			else {
				s[i] = new Stu();
				s[i].index = i;
				s[i].id = temp;
				s[i].name = scan.next();
				s[i].score = scan.nextInt();
			}
		}
		Arrays.sort(s, 0, i, new mycmp());
		for (int j = 0; j < i; j++) {
			System.out.println(s[j].id + " " + s[j].name + " " + s[j].score);
		}
	}
}
```
