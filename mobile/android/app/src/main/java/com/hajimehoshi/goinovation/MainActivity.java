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

    private EbitenView getEbitenView() {
        return (EbitenView)this.findViewById(R.id.ebitenview);
    }

    @Override
    protected void onPause() {
        super.onPause();
        this.getEbitenView().onPause();
    }

    @Override
    protected void onResume() {
        super.onResume();
        this.getEbitenView().onResume();
    }
}
