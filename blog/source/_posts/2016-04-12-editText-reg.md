---
title: 带正则校验的EditText
date: 2016-04-12 09:56:49
categories:
  - Android
  - 自定义View
tags:
  - 自定义View
---

这是好久之前写的了，现在整理下文章那就写一下吧，直接上演示Demo，如下所示，伸手党们可直接跳到最后去下载源码。

{% asset_img acet.gif [演示动画] %}

原理很简单，实现TextWatcher接口中的afterTextChanged()函数，根据正则表达式返回的结果来动态加载rightDrawable，至于校验成功或者失败的drawable用户在初始化的时候可以自定义，下面给出源码参考，其实现了常用的手机、座机、邮箱、url、汉字、用户名的校验，如果你觉得还不满足的话，最后还提供了用户自定义正则的校验，这样总满意了吧，下面给出其代码供参考。
<!-- more -->
``` java
import android.content.Context;
import android.graphics.drawable.Drawable;
import android.text.Editable;
import android.text.TextWatcher;
import android.util.AttributeSet;
import android.widget.EditText;

/**
 * Created by cmj on 2016/4/11.
 */
public class AutoCheckEditText extends EditText implements TextWatcher {
    //手机校验类型
    public static final int TYPE_OF_MOBILE = 0xb0;
    //座机校验类型
    public static final int TYPE_OF_TEL = 0xb1;
    //邮箱校验类型
    public static final int TYPE_OF_EMAIL = 0xb2;
    //url校验类型
    public static final int TYPE_OF_URL = 0xb3;
    //汉字校验类型
    public static final int TYPE_OF_CHZ = 0xb4;
    //用户名校验类型
    public static final int TYPE_OF_USERNAME = 0xb5;

    //用户自定义
    public static final int TYPE_OF_USER_DEFINE = 0xbb;
    private int type;
    private Drawable successDrawable;
    private Drawable unsuccessDrawable;
    private String userRegx;

    public AutoCheckEditText(Context context) {
        super(context);
    }

    public AutoCheckEditText(Context context, AttributeSet attrs) {
        super(context, attrs);
    }

    public AutoCheckEditText(Context context, AttributeSet attrs, int defStyleAttr) {
        super(context, attrs, defStyleAttr);
    }


    /**
     * @param typ       要校验的类型
     * @param success   匹配成功时的图标
     * @param unsuccess 匹配失败时的图标
     */
    public void creatCheck(int typ, Drawable success, Drawable unsuccess) {
        type = typ;
        successDrawable = success;
        successDrawable.setBounds(0, 0, successDrawable.getMinimumWidth(), successDrawable.getMinimumHeight());
        unsuccessDrawable = unsuccess;
        unsuccessDrawable.setBounds(0, 0, unsuccessDrawable.getMinimumWidth(), unsuccessDrawable.getMinimumHeight());
        this.addTextChangedListener(this);
    }

    /**
     * @param typ       要校验的类型
     * @param success   匹配成功时的图标
     * @param unsuccess 匹配失败时的图标
     */
    public void creatCheck(int typ, Drawable success, Drawable unsuccess, String userRegex) {
        creatCheck(typ, success, unsuccess);
        this.userRegx = userRegex;
    }


    @Override
    public void beforeTextChanged(CharSequence s, int start, int count, int after) {
        System.out.println("a");
    }

    @Override
    public void afterTextChanged(Editable s) {
        if (s != null && s.length() > 0) {
            boolean isMatch;
            if (type == TYPE_OF_MOBILE) {
                isMatch = Check.isMobile(s.toString());
            } else if (type == TYPE_OF_TEL) {
                isMatch = Check.isTel(s.toString());
            } else if (type == TYPE_OF_EMAIL) {
                isMatch = Check.isEmail(s.toString());
            } else if (type == TYPE_OF_URL) {
                isMatch = Check.isURL(s.toString());
            } else if (type == TYPE_OF_CHZ) {
                isMatch = Check.isChz(s.toString());
            } else if (type == TYPE_OF_USERNAME) {
                isMatch = Check.isUsername(s.toString());
            } else if (type == TYPE_OF_USER_DEFINE) {
                isMatch = Check.isMatch(userRegx, s.toString());
            } else {
                return;
            }
            if (isMatch) {
                setCompoundDrawables(null, null, successDrawable, null);
            } else {
                setCompoundDrawables(null, null, unsuccessDrawable, null);
            }
        } else {
            setCompoundDrawables(null, null, null, null);
        }
    }

    @Override
    public void onTextChanged(CharSequence s, int start, int before, int count) {
        System.out.println();
    }
}

```

最后献上源码地址：**[AutoCheckEditTextDemo](https://github.com/Blankj/AutoCheckEditTextDemo)**，欢迎大家star和fork。