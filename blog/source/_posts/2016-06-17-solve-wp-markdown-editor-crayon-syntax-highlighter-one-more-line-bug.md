---
title: 解决WP Markdown Editor和Crayon Syntax Highlighter一起使用代码块无故多一行
date: 2016-06-17 15:06:42
categories:
  - WordPress
tags:
  - WordPress
---

　　当用WP Markdown Editor写完文章发布后，如果存在代码块的话，在前台由Crayon Syntax Highlighter渲染出来会存在如下图问题。

{% asset_img oneMoreLine.png [oneMoreLine] %}

　　也就是代码块无故多一行出来，影响美观，由审查元素可以看到多的是一个转义的空格&nbsp;

　　处理了大半天，终于解决了，是目前来说最好的解决办法了，方法如下：

　　在主题的footer.php最后加上如下JS代码即可，代码一目了然是在做什么。
<!--more-->
``` js
<script> 
var nodes = document.querySelectorAll('.crayon-pre');  
for (var i = 0; i < nodes.length; i++) {
    var lastnode = nodes[i].lastChild;  
    if (lastnode.innerHTML == ' ') {
        nodes[i].removeChild(lastnode)
        var parent = nodes[i].parentNode.parentNode.parentNode.firstChild.childNodes[1].childNodes[1];
        var parentLastNode = parent.lastChild;
        parent.removeChild(parentLastNode);
    }
}
</script>
```

　　再清一下缓存去前台探探路，如下图所示。

{% asset_img normalLine.png [normalLine] %}

　　大功告成。

* * *

原文地址[解决WP Markdown Editor和Crayon Syntax Highlighter一起使用代码块无故多一行][passage_url]

我的自媒体博客[blankj小站](http://blankj.com/)，欢迎来逛逛。

[passage_url]: http://blankj.com/2016/06/17/解决WP-Markdown-Editor和Crayon-Syntax-Highlighter一起使用代码块无故多一行/