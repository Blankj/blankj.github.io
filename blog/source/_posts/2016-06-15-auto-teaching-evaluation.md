---
title: 正方教务系统自动评教
date: 2016-06-15 19:30:35
categories:
  - JS
  - 奇淫技巧
tags:
  - 奇淫技巧
---

　　期末到了，又到了评教的时刻，是时候展示曾经学的JS功底了，这里以推荐谷歌浏览器去评教，首先到进入到评教页面，点击老师后弹出新的页面，如下图所示：

{% asset_img pingjiao.png [评教页面] %}
<!--more-->
　　打开浏览器控制台（Ctrl + Shift + C），粘贴下方的JS代码在控制台中，点击回车即可，具体情况如下图所示：

``` js
function autoPJXT_YONIIIIII_ZJCM() {
    for(var j = 0; j < 13; ++j) {
        var sel = document.getElementsByTagName('select')[j];
        for (var i = 0; i < sel.length; i++) {
            if (sel[i].value == "86-101") {
                sel[i].selected = true;
                document.getElementsByTagName('input')[j+1].value=100.9;
            }
        }
        document.getElementById("txt_pjxx").value='老师备课充分，授课重点突出。';
    }
    document.getElementById("Button1").click();
}autoPJXT_YONIIIIII_ZJCM();
```

{% asset_img js_console.png [控制台操作] %}

　　关闭控制台后，返回到页面中，可以发现结果都已自动写好，并提交完毕，turn to next one

{% asset_img pingjiaofinish.png [运行结果] %}

　　是不是省事多了哈。