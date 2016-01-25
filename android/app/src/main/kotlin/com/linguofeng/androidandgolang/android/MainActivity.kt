package com.linguofeng.androidandgolang.android

import android.os.Bundle
import android.support.v7.app.AppCompatActivity
import android.util.Log
import go.golang.Golang

class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)

        val hello = Golang.Hello("linguofeng")

        Log.d("MainActivity", hello)

        Golang.SendStr(hello, object : Golang.Callback.Stub() {
            override fun CallByGo(str: String?) {
                Log.d("MainActivity", str)
            }
        })
    }
}
