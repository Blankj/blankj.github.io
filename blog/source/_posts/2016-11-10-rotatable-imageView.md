---
title: 可旋转的ImageView
date: 2016-11-10 17:09:18
categories:
  - Android
  - 自定义View
tags:
  - 自定义View
---

话不多少，先上效果图过一下瘾。

{% asset_img demo.gif [演示动画] %}
<!-- more -->
经我网上一搜，没有很好的实例来单独实现以上效果，所以只能老司机自己出马了。为了实现以上效果，当然离不开我们对`onTouchEvent`事件的处理了，老司机们应该都懂的，开始开车，嘟嘟~

{% asset_img 开车.gif [开车] %}

下面来说一下思路，首先上一张图，让我更方便地来为乘客们讲解。

{% asset_img 旋转解析图.jpg [旋转解析图] %}

好了，重要的都已经在图中说了，下面只需要一个回调便可在外部获取到旋转的角度，献上相关重要代码，相信根据我的注释，大家也都应该明白了。

``` java
import android.content.Context;
import android.graphics.Bitmap;
import android.graphics.Canvas;
import android.graphics.Matrix;
import android.graphics.drawable.BitmapDrawable;
import android.util.AttributeSet;
import android.util.Log;
import android.view.MotionEvent;
import android.widget.ImageView;

/**
 * <pre>
 *     author: Blankj
 *     blog  : http://blankj.com
 *     time  : 2016/11/8
 *     desc  :
 * </pre>
 */
public class RotateImageView extends ImageView {

    // 点击在view中所在坐标
    private float x, y;
    // view的中心点坐标
    private float ox, oy;
    // view的宽、高
    private int w, h;
    private Matrix matrix;
    // ImageView的资源图片
    Bitmap mBitmap;
    // 旋转过的角度
    public  float   totalDegree = 0;
    // 是否初始化过
    private boolean isInit      = false;

    // 旋转角度改变的监听器
    private OnRotationChangeListener mListener;

    public void setOnRotationChangeListener(OnRotationChangeListener listener) {
        mListener = listener;
    }

    public RotateImageView(Context context) {
        this(context, null);
    }

    public RotateImageView(Context context, AttributeSet attrs) {
        this(context, attrs, 0);
    }

    public RotateImageView(Context context, AttributeSet attrs, int defStyleAttr) {
        super(context, attrs, defStyleAttr);
    }

    @Override
    protected void onMeasure(int widthMeasureSpec, int heightMeasureSpec) {
        super.onMeasure(widthMeasureSpec, heightMeasureSpec);
        // 已初始化完毕的话那就直接滚回老家去吧
        if (isInit) return;
        int measureWidth = getWidth();
        int measureHeight = getHeight();
        Log.d("blankj", "onMeasure: w: " + measureWidth + ", h: " + measureHeight);
        // 如果没有初始化过，就初始化
        if (measureWidth != 0 && measureHeight != 0 && !isInit) {
            isInit = true;
            mBitmap = ((BitmapDrawable) this.getDrawable()).getBitmap();
            w = measureWidth;
            h = measureHeight;
            ox = w >> 1;
            oy = h >> 1;
            matrix = new Matrix();
            // 获取适配于ImageView的bitmap
            mBitmap = Bitmap.createScaledBitmap(mBitmap, w, h, true);
        }
    }

    @Override
    protected void onDraw(Canvas canvas) {
        Log.d("blankj", "onDraw");
        // 判空以防崩溃
        if (mBitmap == null) return;
        canvas.save();
        canvas.drawBitmap(mBitmap, matrix, null);
        canvas.restore();
    }

    @Override
    public boolean onTouchEvent(MotionEvent event) {
        switch (event.getAction()) {
            case MotionEvent.ACTION_DOWN:
                x = event.getX();
                y = event.getY();
                break;
            case MotionEvent.ACTION_MOVE:
                float nowX = event.getX();
                float nowY = event.getY();

                // 计算三边的平方
                float ab2 = (x - nowX) * (x - nowX) + (y - nowY) * (y - nowY);
                float oa2 = (x - ox) * (x - ox) + (y - oy) * (y - oy);
                float ob2 = (nowX - ox) * (nowX - ox) + (nowY - oy) * (nowY - oy);

                // 根据两向量的叉乘来判断顺逆时针
                boolean isClockwise = ((x - ox) * (nowY - oy) - (y - oy) * (nowX - ox)) > 0;

                // 根据余弦定理计算旋转角的余弦值
                double cosDegree = (oa2 + ob2 - ab2) / (2 * Math.sqrt(oa2) * Math.sqrt(ob2));

                // 异常处理，因为算出来会有误差绝对值可能会超过一，所以需要处理一下
                if (cosDegree > 1) {
                    cosDegree = 1;
                } else if (cosDegree < -1) {
                    cosDegree = -1;
                }

                // 计算弧度
                double radian = Math.acos(cosDegree);

                // 计算旋转过的角度，顺时针为正，逆时针为负
                float degree = (float) (isClockwise ? Math.toDegrees(radian) : -Math.toDegrees(radian));

                // 累加角度
                totalDegree += degree;
                matrix.setRotate(totalDegree, ox, oy);

                // 更新触摸点
                x = nowX;
                y = nowY;

                // 回调把角度抛出
                if (mListener != null) {
                    mListener.getRotation((int) totalDegree);
                }

                invalidate();
                break;
            case MotionEvent.ACTION_UP:
                // 如果图片需要复原原来角度，调用下方代码
//                matrix.setRotate(totalDegree, ox, oy);
//                invalidate();
                break;
        }
        return true;
    }

    public void reset() {
        totalDegree = 0;
        matrix.setRotate(totalDegree, ox, oy);
        invalidate();
    }

    public interface OnRotationChangeListener {
        void getRotation(int degree);
    }
}
```

好了，乘客们，终点站已到，谢谢你们旅途中的陪伴，老司机马上要下班了，啊哈哈哈，还是那句话：“你不买东西也是没钱，买了东西也是没钱，那就说明买东西它不要钱啊”，祝大家双十一购物愉快。

{% asset_img good.jpg %}

最后献上源码地址：**[GetRotateDegree](https://github.com/Blankj/GetRotateDegree)**，欢迎大家star和fork。