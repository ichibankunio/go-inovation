package com.hajimehoshi.goinovation;

import android.content.Context;
import android.util.AttributeSet;
import android.util.Log;

import com.hajimehoshi.goinovation.mobile.EbitenView;

class EbitenViewWithErrorHandling extends EbitenView {
    public EbitenViewWithErrorHandling(Context context) {
        super(context);
    }

    public EbitenViewWithErrorHandling(Context context, AttributeSet attributeSet) {
        super(context, attributeSet);
    }

    @Override
    protected void onErrorOnGameUpdate(Exception e) {
        // You can define your own error handling e.g., using Crashlytics.
        Log.e("Inovation Error!", e.toString());
    }
}
