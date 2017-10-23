---
title: Git的简单使用
date: 2016-03-24 11:22:51
categories:
  - Android
tags:
  - Android Tips
---

```
$ mkdir learngit
$ cd learngit
$ pwd
/Users/michael/learngit

//初始化一个Git仓库
$ git init  

//添加文件到Git仓库，分两步：
$ git add readme.txt
$ git commit -m "..."

//查看仓库当前的状态
$ git status

//查看difference
$ git diff readme.txt 

//查看历史记录
$ git log

//信息太多可用如下
$ git log --pretty=oneline

//用HEAD表示当前版本，上一个版本就是HEAD^，上上一个版本就是HEAD^^
//当然往上100个版本写100个^比较容易数不过来，所以写成HEAD~100。
//回退到上一个版本
$ git reset --hard HEAD^

//查看命令历史
$ git reflog

//回退到id是1234..的版本
$ git reset --hard 1234..

//回到最近一次git commit或git add时的状态
$ git checkout -- readme.txt

//把暂存区的修改撤销掉（unstage），重新放回工作区
$ git reset HEAD readme.txt

//删除一个文件
$ git rm test.txt

//关联一个远程库
$ git remote add origin http://192.168.0.62:82/cmj/learnGit.git

//推送master分支的所有内容，远程的话commit之后，最后还需要push到远程库，远程库才会发生修改
$ git push -u origin master

//推送最新修改
$ git push origin master

//用http协议克隆一个仓库
$ git clone http://192.168.0.62:82/cmj/learnGit.git

//我们创建dev分支，然后切换到dev分支
$ git checkout -b dev
//等同于以下两句
{
//创建分支
$ git branch dev
//切换分支到dev
$ git checkout dev
}

//查看当前分支，前面带*的就是当前分支
$ git branch

//分支某合并到当前分支
$ git merge dev

//删除分支
$ git branch -d dev
```