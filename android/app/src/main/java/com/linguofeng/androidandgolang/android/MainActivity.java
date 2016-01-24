package com.linguofeng.androidandgolang.android;

import android.os.Bundle;
import android.support.v7.app.AppCompatActivity;
import android.util.Log;

import go.Seq;
import go.android.Android;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        String hello = Android.Hello("linguofeng");

        Log.d("MainActivity", hello);

        Android.CallJava(new Android.Callback.Stub() {
            @Override
            public void CallJava() {
                Log.d("MainActivity", "go call java");
            }
        });
    }
}
