package com.hajimehoshi.goinovation;

import android.content.Context;
import android.opengl.GLSurfaceView;
import android.util.Log;
import android.util.AttributeSet;
import android.view.MotionEvent;

import javax.microedition.khronos.egl.EGLConfig;
import javax.microedition.khronos.opengles.GL10;

import com.hajimehoshi.goinovation.ebitenmobileview.Ebitenmobileview;
import com.hajimehoshi.goinovation.ebitenmobileview.ViewRectSetter;

public class EbitenSurfaceView extends GLSurfaceView {

    private class EbitenRenderer implements Renderer {

        private boolean mErrored;

        @Override
        public void onDrawFrame(GL10 gl) {
            if (mErrored) {
                return;
            }
            try {
                Ebitenmobileview.update();
            } catch (Exception e) {
                Log.e("Go Error", e.toString());
                mErrored = true;
            }
        }

        @Override
        public void onSurfaceCreated(GL10 gl, EGLConfig config) {
        }

        @Override
        public void onSurfaceChanged(GL10 gl, int width, int height) {
        }
    }

    private double mDeviceScale = 0.0;
    private boolean mRunning = false;

    public EbitenSurfaceView(Context context) {
        super(context);
        initialize();
    }

    public EbitenSurfaceView(Context context, AttributeSet attrs) {
        super(context, attrs);
        initialize();
    }

    private void initialize() {
        setEGLContextClientVersion(2);
        setEGLConfigChooser(8, 8, 8, 8, 0, 0);
        setRenderer(new EbitenRenderer());
    }

    private double deviceScale() {
        if (mDeviceScale == 0.0) {
            mDeviceScale = getResources().getDisplayMetrics().density;
        }
        return mDeviceScale;
    }

    private double pxToDp(double x) {
        return x / deviceScale();
    }

    private double dpToPx(double x) {
        return x * deviceScale();
    }

    @Override
    public void onLayout(boolean changed, int left, int top, int right, int bottom) {
        super.onLayout(changed, left, top, right, bottom);

        int width = (int)Math.floor(pxToDp(right - left));
        int height = (int)Math.floor(pxToDp(bottom - top));
        Ebitenmobileview.layout(width, height, new ViewRectSetter() {
            @Override
            public void setViewRect(long x, long y, long width, long height) {
                int oldWidth = getLayoutParams().width;
                int oldHeight = getLayoutParams().height;
                int newWidth = (int)Math.ceil(dpToPx(width));
                int newHeight = (int)Math.ceil(dpToPx(height));
                if (oldWidth == newWidth && oldHeight == newHeight) {
                    return;
                }
                getLayoutParams().width = newWidth;
                getLayoutParams().height = newHeight;
                post(new Runnable() {
                    @Override
                    public void run() {
                        requestLayout();
                    }
                });
            }
        });
    }

    @Override
    public boolean onTouchEvent(MotionEvent e) {
        for (int i = 0; i < e.getPointerCount(); i++) {
            int id = e.getPointerId(i);
            int x = (int)e.getX(i);
            int y = (int)e.getY(i);
            Ebitenmobileview.updateTouchesOnAndroid(e.getActionMasked(), id, (int)pxToDp(x), (int)pxToDp(y));
        }
        return true;
    }
}
