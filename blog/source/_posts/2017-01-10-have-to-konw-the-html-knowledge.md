---
title: 你不得不了解的HTML知识
date: 2016-12-27 18:02:07
categories:
  - JS
  - HTML
tags:
  - HTML
---

> 这次要细说的只是块级元素和行内元素的内容，如有不对，请轻喷。

## 块级元素和行内元素
默认情况下块级元素会始终占居一行，而行内元素并不会。除了 table 元素的 display 属性比较特殊以外，基本上所有的 HTML 元素的 display 的属性值要么是 block，要么是 inline。无论你想了解哪个 HTML 元素，第一个要问的问题就是：它是块级元素还是行内元素，然后在编写标记的时候预想到这个元素在初始状态下是如何定位的，这样才能进一步想好将来怎么用 CSS 重新定位它，因为块级元素和行内元素在定位上有很大的区别，后面会详细说明。

块级元素盒子（一个很重要的概念————盒模型）会扩展到与父元素同宽，这也是为什么块级元素会占居一行的原因了，因为所有块级元素的父元素都是 body，而它的默认宽度就是浏览器的视口大小，所以默认情况下块级元素的宽度也和浏览器的视口一样宽，这样以来，一个块级元素旁边也就没有空间来容纳另一个块级元素了。

相比于块级元素会扩展到与父元素同宽，然而行内元素的行为却是恰恰相反，它会尽量的收缩包裹其内容，这也就是为什么几个行内元素会并排显示在一行直到它们排满一行才会另起一行，而每个块级元素会直接另起一行的原因了。

<!-- more -->

下表列出了一些常见的块级元素和行内元素：

| 块级元素 | 行内元素  |
|  :---: |  :---: |
| div| span|
| form| a|
| table| img|
| header| label|
| aside| input|
| section| select|
| article| textarea|
| figure| br|
| figcaption| i|
| h1~h6| em|
| nav| strong|
| p| small|
| pre| button|
| blockqoute| sub|
| canvas| sup|
| ol, ul, dl| code|

之前提到过无论你想了解哪个 HTML 元素，第一个要问的问题就是：它是块级元素还是行内元素，因为它们在盒模型上的表现有很大的不同，不过在了解它们的不同之前我们还得先知道另外一个概念————[替换元素](http://www.w3.org/TR/html5/rendering.html#replaced-elements)和[非替换元素](http://www.w3.org/TR/html5/rendering.html#non-replaced-elements)，其中替换元素就是指浏览器是根据元素的属性来判断具体要显示的内容的元素，比如 `img` 标签，浏览器是根据其 `src` 的属性值来读取这个元素所包含的内容的，常见的替换元素还有 `input` 、`textarea`、 `select`、 `object`、 `iframe` 和 `video` 等等，这些元素都有一个共同的特点，就是浏览器并不直接显示其内容，而是通过其某个属性的值来显示具体的内容，比如浏览器会根据 `input` 中的 `type` 的属性值来判断到底应该显示单选按钮还是多选按钮亦或是文本输入框。而对于非替换元素，比如 `p`、`label` 元素等等，浏览器这是直接显示元素所包含的内容。看到这里你应该大概的知道了什么是替换元素和非替换元素了。

对着两个概念有了大概的了解后就可以对 `block` 和 `inline` 在盒模型上的表现差异进行了解了，首先是 `margin` ，[W3C](http://www.w3.org/TR/CSS2/box.html#margin-properties) 对其所支持了元素对象是这么定义的：

> Applies to: all elements except elements with table display types other than table-caption, table and inline-table

我的理解就是所有元素都支持 `margin` 除了 `display` 属性值为 `table-caption` 和 `table-inline` 以外的所有表格显示类型比如`table-row-group`、 `table-cell`、 `table-row` 和 `table-header-group`等等,但是为了验证我的理解，我发现 `display` 属性值为 `table` 的元素也支持，可能是我对原文标准的理解有误。但还有一个要特别注意的是 `margin-top` 和 `margin-bottom` 两个属性比较特殊，它们对非替换行内元素没有效果，下面是 W3C 上对于 `margin-top` 和`margin-bottom` 支持对象的介绍：

> Applies to: all elements except elements with table display types other than table-caption, table and inline-table
These properties have no effect on non-replaced inline elements.

前面一句和之前对 `margin` 的描述是一样的，这毫无疑问，下面这句话的意思是这些（ `margin-top` 和 `margin-bottom` ）属性对非替换行内元素没有效果比如 `a` 和`span`，注意这里是**非替换行内元素**而不单单是非替换元素或者是行内元素。比如 `img` 就是一个行内元素， `margin-top` 和 `margin-bottom` 对它是有效果的，因为它是一个替换元素而不是非替换元素，所以对于「 `margin-top` 和 `margin-bottom` 对行内元素没有效果」这种说法是不对的。

而对于 `padding` 的支持对象，W3C 是这么描述的：

> all elements except table-row-group, table-header-group, table-footer-group, table-row, table-column-group and table-column

上面这句话的意思是除了表格显示类型为 `table-row-group`、 `table-header-group`、 `table-footer-group`、 `table-row`,` table-column-group` 和 `table-column` 的元素不支持，其他所有的元素都支持。

但这里有些特殊情况需要注意的是，对行内元素比如 `span` 和 `img` 设置左右内边距的效果是可见可，但是对行内元素设置上下内边距在有些情况下是不可见的，这些情况又要分为是否为替换元素和是否设置了背景色，为了能更直观的了解这些概念，我在这里做了个表格：

|  padding-top和padding-bottom 对于行内元素是否可见|替换元素（e.g: input）| 非替换元素（e.g: span）|
| :---:|:---:|:---:|
|<br>设置背景色|可见<br>影响行高<br>会撑开父元素|可见<br>不影响行高<br>不会撑开父元素|
|<br>没有设置背景色|可见<br>影响行高<br>会撑开父元素|不可见<br>不影响行高<br>不会撑开父元素|

所以对于「 `padding-top` 和 `padding-bottom` 对行内元素没有效果」这种说法也是不对的，因为它们只是对于没有设置背景色的行内非替换元素效果不可见而已，而对于行内替换元素来说，不管是否设置了背景色都是有效果了，并且会把父元素撑开。

说了这么多 `block` 和 `inline` 的区别，其实除了这两个常见的 `display` 属性以外还有一个属性也是非常常见的，那就是 `inline-block` ，没错，这就是前面两种情况的结合体，它既有 `block` 的特性又有 `inline` 的特性，比如把一个 `display` 属性值为 `block` 或者 `inline` 的元素属性值设置成 `inline-block` 后，既可以用只对行内元素有效的 `text-align: center;` 声明对其进行居中以外，还可以用 `padding-top` 和 `padding-bottom` 对元素设置上下内边距而无需对其设置背景色，并且能把父元素撑开。

不要问我为何知道这么多，我只献上[传送门](http://www.w3cschool.cn/)
