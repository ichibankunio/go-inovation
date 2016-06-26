package com.hajimehoshi.goinovation;

import android.content.Context;
import android.opengl.GLSurfaceView;
import android.util.Log;
import android.util.AttributeSet;
import android.view.MotionEvent;
import android.view.View;

import javax.microedition.khronos.egl.EGLConfig;
import javax.microedition.khronos.opengles.GL10;

import com.hajimehoshi.goinovation.go.Inovation;

public class EbitenGLSurfaceView extends GLSurfaceView {

    private class EbitenRenderer implements Renderer {

        private boolean mErrored;

        @Override
        public void onDrawFrame(GL10 gl) {
            if (mErrored) {
                return;
            }
            try {
                Inovation.Render();
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

    public EbitenGLSurfaceView(Context context) {
        super(context);
        initialize();
    }

    public EbitenGLSurfaceView(Context context, AttributeSet attrs) {
        super(context, attrs);
        initialize();
    }

    private void initialize() {
        setEGLContextClientVersion(2);
        setEGLConfigChooser(8, 8, 8, 8, 0, 0);
        setRenderer(new EbitenRenderer());
    }

    public double getScale() {
        View parent = (View)getParent();
        return Math.max(1,
                Math.min(parent.getWidth() / (double)Inovation.ScreenWidth,
                        parent.getHeight() / (double)Inovation.ScreenHeight));
    }

    @Override
    public void onLayout(boolean changed, int left, int top, int right, int bottom) {
        super.onLayout(changed, left, top, right, bottom);
        double scale = getScale();
        getLayoutParams().width = (int)(Inovation.ScreenWidth * scale);
        getLayoutParams().height = (int)(Inovation.ScreenHeight * scale);
        try {
            if (!Inovation.IsRunning()) {
                Inovation.Start(scale);
            }
        } catch (Exception e) {
            Log.e("Go Error", e.toString());
        }
    }

    @Override
    public boolean onTouchEvent(MotionEvent e) {
        for (int i = 0; i < e.getPointerCount(); i++) {
            int id = e.getPointerId(i);
            int x = (int) e.getX(i);
            int y = (int) e.getY(i);
            Inovation.UpdateTouchesOnAndroid(e.getActionMasked(), id, x, y);
        }
        return true;
    }
}
