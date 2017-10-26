# Blankj's Blog

## About Blog

### 主题[next](https://github.com/iissnan/hexo-theme-next)

### 评论[gitalk](https://github.com/gitalk/gitalk)

需要 `Github Application`，如果没有 [点击这里申请](https://github.com/settings/applications/new)，`Authorization callback URL` 填写你主页地址，比如我的就是`https://blankj.com`，其他都随意。

1. 首先创建 `Gitalk` 的 `swig` 文件，放在 `themes/next/layout/_third-party/comments` 文件夹下，命名为 `gitalk.swig` 。内容如下

```js
{% if page.comments && theme.gitalk.enable %}
  <link rel="stylesheet" href="https://unpkg.com/gitalk/dist/gitalk.css">

  <script src="https://unpkg.com/gitalk/dist/gitalk.min.js"></script>
   <script type="text/javascript">
		var gitalk = new Gitalk({
		  clientID: '{{ theme.gitalk.ClientID }}',
		  clientSecret: '{{ theme.gitalk.ClientSecret }}',
		  repo: '{{ theme.gitalk.repo }}',
		  owner: '{{ theme.gitalk.githubID }}',
		  admin: ['{{ theme.gitalk.adminUser }}'],
		  id: location.pathname,
		  distractionFreeMode: '{{ theme.gitalk.distractionFreeMode }}'
		})
		gitalk.render('gitalk-container')           
       </script>
{% endif %}
```

2. 在主题文件 `themes/next/layout/_third-party/comments/index.swig` 中引入刚刚添加的文件。

```js
{% include 'gitalk.swig' %}
```

3. 在 `themes/next/layout/_partials/comments.swig` 文件中找到倒数第二个 `{% endif %}` 语句，在前面插入如下代码:

```js
{% elseif theme.gitalk.enable %}
  <div id="gitalk-container"></div>
```

4. 在 `themes/next/_config.xml` 文件中引入 `Gitalk`

```
gitalk:
  enable: true
  githubID: Blankj
  repo: xxx
  ClientID: xxx
  ClientSecret: xxx
  adminUser: Blankj
  distractionFreeMode: true
```

其中 `githubID` 是你的 `Github` 用户名，`repo` 是你用来存放评论 `issue` 的仓库，比如我的就是[blog-comment](https://github.com/Blankj/blog-comment)，那么我就写 `blog-comment` 即可，`ClientID` 和 `ClientSecret` 就是你之前申请 `Github Application` 可以获取到，`adminUser` 和 `githubID` 一样即可，`distractionFreeMode` 是评论时遮照效果的开关。

### Logo字体

在 `themes/next/source/css/_custom/custom.styl` 中设置如下即可。

```css
@font-face {
    font-family: Taken;
    src: url('/fonts/Blankj.ttf');
}

.site-title {
    font-size: 40px !important;
	font-family: 'Taken' !important;
}
```

其中字体文件在 `themes/next/source/fonts` 目录下，里面有个 `.gitkeep` 的隐藏文件，打开写入你要保留的字体文件，比如我的是就是写入 `Blankj.ttf`，具体字库自己从网上下载即可。

### 圆形头像

修改 `themes/next/source/css/_common/components/sidebar/sidebar-author.styl` 为下方代码。

```css
.site-author-image {
  display: block;
  margin: 0 auto;
  padding: $site-author-image-padding;
  max-width: $site-author-image-width;
  height: $site-author-image-height;
  border: $site-author-image-border-width solid $site-author-image-border-color;
  /* 头像圆形 */
  border-radius: 80px;
  -webkit-border-radius: 80px;
  -moz-border-radius: 80px;
  box-shadow: inset 0 -1px 0 #333sf;
  /* 设置循环动画 [animation: (play)动画名称 (2s)动画播放时长单位秒或微秒 (ase-out)动画播放的速度曲线为以低速结束
    (1s)等待1秒然后开始动画 (1)动画播放次数(infinite为循环播放) ]*/
  -webkit-animation: play 2s ease-out 1s 1;
  -moz-animation: play 2s ease-out 1s 1;
  animation: play 2s ease-out 1s 1;
  /* 鼠标经过头像旋转360度 */
  -webkit-transition: -webkit-transform 1.5s ease-out;
  -moz-transition: -moz-transform 1.5s ease-out;
  transition: transform 1.5s ease-out;
}
img:hover {
  /* 鼠标经过停止头像旋转
  -webkit-animation-play-state:paused;
  animation-play-state:paused;*/
  /* 鼠标经过头像旋转360度 */
  -webkit-transform: rotateZ(360deg);
  -moz-transform: rotateZ(360deg);
  transform: rotateZ(360deg);
}
/* Z 轴旋转动画 */
@-webkit-keyframes play {
  0% {
    -webkit-transform: rotateZ(0deg);
  }
  100% {
    -webkit-transform: rotateZ(-360deg);
  }
}
@-moz-keyframes play {
  0% {
    -moz-transform: rotateZ(0deg);
  }
  100% {
    -moz-transform: rotateZ(-360deg);
  }
}
@keyframes play {
  0% {
    transform: rotateZ(0deg);
  }
  100% {
    transform: rotateZ(-360deg);
  }
}
.site-author-name {
  margin: $site-author-name-margin;
  text-align: $site-author-name-align;
  color: $site-author-name-color;
  font-weight: $site-author-name-weight;
}
.site-description {
  margin-top: $site-description-margin-top;
  text-align: $site-description-align;
  font-size: $site-description-font-size;
  color: $site-description-color;
}
```
