package com.hajimehoshi.goinovation;

import android.content.Context;
import android.opengl.GLSurfaceView;
import android.util.Log;
import android.util.AttributeSet;
import android.view.MotionEvent;

import javax.microedition.khronos.egl.EGLConfig;
import javax.microedition.khronos.opengles.GL10;

import com.hajimehoshi.goinovation.ebitenmobileview.Ebitenmobileview;

class EbitenSurfaceView extends GLSurfaceView {

    private class EbitenRenderer implements Renderer {

        private boolean errored_ = false;

        @Override
        public void onDrawFrame(GL10 gl) {
            if (errored_) {
                return;
            }
            try {
                Ebitenmobileview.update();
            } catch (Exception e) {
                for (String line : e.toString().split("\\n")) {
                    Log.e("Go Error", line);
                }
                errored_ = true;
            }
        }

        @Override
        public void onSurfaceCreated(GL10 gl, EGLConfig config) {
        }

        @Override
        public void onSurfaceChanged(GL10 gl, int width, int height) {
        }
    }

    private double deviceScale_ = 0.0;

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

    private double getDeviceScale() {
        if (deviceScale_ == 0.0) {
            deviceScale_ = getResources().getDisplayMetrics().density;
        }
        return deviceScale_;
    }

    private double pxToDp(double x) {
        return x / getDeviceScale();
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
