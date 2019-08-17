package com.hajimehoshi.goinovation;

import android.content.Context;
import android.util.AttributeSet;
import android.view.ViewGroup;

import com.hajimehoshi.goinovation.ebitenmobileview.Ebitenmobileview;
import com.hajimehoshi.goinovation.ebitenmobileview.ViewRectSetter;

public class EbitenView extends ViewGroup {
    private double getDeviceScale() {
        if (deviceScale_ == 0.0) {
            deviceScale_ = getResources().getDisplayMetrics().density;
        }
        return deviceScale_;
    }

    private double pxToDp(double x) {
        return x / getDeviceScale();
    }

    private double dpToPx(double x) {
        return x * getDeviceScale();
    }

    private double deviceScale_ = 0.0;

    public EbitenView(Context context) {
        super(context);
        ebitenSurfaceView_ = new EbitenSurfaceView(context);
    }

    public EbitenView(Context context, AttributeSet attrs) {
        super(context, attrs);
        ebitenSurfaceView_ = new EbitenSurfaceView(context, attrs);
    }

    @Override
    protected void onLayout(boolean changed, int left, int top, int right, int bottom) {
        if (!initialized_) {
            LayoutParams params = new LayoutParams(LayoutParams.WRAP_CONTENT, LayoutParams.WRAP_CONTENT);
            addView(ebitenSurfaceView_, params);
            initialized_ = true;
        }

        int widthInDp = (int)Math.floor(pxToDp(right - left));
        int heightInDp = (int)Math.floor(pxToDp(bottom - top));
        Ebitenmobileview.layout(widthInDp, heightInDp, new ViewRectSetter() {
            @Override
            public void setViewRect(long xInDp, long yInDp, long widthInDp, long heightInDp) {
                int widthInPx = (int)Math.ceil(dpToPx(widthInDp));
                int heightInPx = (int)Math.ceil(dpToPx(heightInDp));
                int xInPx = (int)Math.ceil(dpToPx(xInDp));
                int yInPx = (int)Math.ceil(dpToPx(yInDp));
                ebitenSurfaceView_.layout(xInPx, yInPx, xInPx + widthInPx, yInPx + heightInPx);
            }
        });
    }

    public void onPause() {
        if (initialized_) {
            ebitenSurfaceView_.onPause();
        }
    }

    public void onResume() {
        if (initialized_) {
            ebitenSurfaceView_.onPause();
        }
    }

    private EbitenSurfaceView ebitenSurfaceView_;
    private boolean initialized_ = false;
}
