---
title: Hay Points
date: 2015-03-24 14:15:26
categories:
  - OJ
tags:
  - 模拟
---

### Description

Each employee of a bureaucracy has a job description - a few paragraphs that describe the responsibilities of the job. The employee's job description, combined with other factors, such as seniority, is used to determine his or her salary.

The Hay Point system frees the Human Resources department from having to make an intelligent judgement as to the value of the employee; the job description is merely scanned for words and phrases that indicate responsibility. In particular, job descriptions that indicate control over a large budget or management over a large number of people yield high Hay Point scores.

You are to implement a simplified Hay Point system. You will be given a Hay Point dictionary and a number of job descriptions. For each job description you are to compute the salary associated with the job, according to the system.

<!-- more -->
### Input

The first line of input contains 2 positive integers: m <= 1000, the number of words in the Hay Point dictionary, and n <= 100, the number of job descriptions. m lines follow; each contains a word (a string of up to 16 lower-case letters) and a dollar value (a real number between 0 and 1,000,000). Following the dictionary are the n job descriptions. Each job description consists of one or more lines of text; for your convenience the text has been converted to lower case and has no characters other than letters, numbers, and spaces. Each job description is terminated by a line containing a period.

### Output

For each job description, output the corresponding salary computed as the sum of the Hay Point values for all words that appear in the description. Words that do not appear in the dictionary have a value of 0.

### Sample Input

<pre>7 2
administer 100000
spending 200000
manage 50000
responsibility 25000
expertise 100
skill 50
money 75000
the incumbent will administer the spending of kindergarden milk money
and exercise responsibility for making change he or she will share
responsibility for the task of managing the money with the assistant
whose skill and expertise shall ensure the successful spending exercise
.
this individual must have the skill to perform a heart transplant and
expertise in rocket science
.</pre>

### Sample Output

<pre>700150
150</pre>

分析：这题如果用映射当然最好不过了，我对STL还不是很熟悉，所以还是用结构体好了，输入的m个单词作为字典，之后输入的n个文本以'.'结尾，如果文本中出现字典中的单词，就加上值，最后输出。

``` cpp
#include<iostream>
#include<algorithm>
#include<vector>
#include<cctype>
#include<cmath>
#include<cstring>
#include<cstdio>
using namespace std; 
struct dic
{
	char str[17];
	int val;
}d[1001];
void solve();
int main()
{
	solve();
	return 0;
}
void solve()
{   
	char temp[17];
	int m,n,i,j,ans;
	cin>>m>>n;
	for(i=0;i<m;i++)
		cin>>d[i].str>>d[i].val;
	for(j=0;j<n;j++)
	{
		ans=0;
		while(cin>>temp&amp;&amp;temp[0]!=&#39;.&#39;)
		{
			for(i=0;i<m;i++)
			{
				if(strcmp(temp,d[i].str)==0)
				{
					ans+=d[i].val;
					break;
				}
			}
		}
		cout<<ans<<endl;
	}
}
```