package com.hajimehoshi.goinovation;

import android.opengl.GLSurfaceView;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;

import go.Seq;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        Seq.setContext(getApplicationContext());
    }

    private GLSurfaceView glSurfaceView() {
        return (GLSurfaceView)this.findViewById(R.id.glview);
    }

    @Override
    protected void onPause() {
        super.onPause();
        this.glSurfaceView().onPause();
    }

    @Override
    protected void onResume() {
        super.onResume();
        this.glSurfaceView().onResume();
    }
}
