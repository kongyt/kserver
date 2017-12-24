using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public delegate void TimerEventHandler(int times);

// 定时器类
public class Timer{

    private float during;
    private float leftTime;
    private int repeat;
    private int times;

    private event TimerEventHandler onTimer;

    private static List<Timer> timerList = new List<Timer>();

    public Timer() {

    }

    public Timer(float during, TimerEventHandler handler) {
        Init(during, 1, true, handler);
    }


    public Timer(float during, int repeat, TimerEventHandler handler) {
        Init(during, repeat, true, handler);
    }


    public Timer(float during, int repeat, bool firstDelay, TimerEventHandler handler) {
        Init(during, repeat, firstDelay, handler);
    }


    // 定时器初始化函数
    // @param during        定时器间隔时间
    // @param repeat        定时器触发次数, 如果repeat == -1 则无限次触发
    // @param firstDelay    第一次触发是否延迟
    // @param timerCallback 定时器触发回调函数
    public void Init(float during, int repeat, bool firstDelay, TimerEventHandler handler) {
        if (repeat == 0){
            return;
        }

        this.during = during;
        this.repeat = repeat;
        
        this.leftTime = during;
        this.times = 0;
        this.onTimer = handler;
        timerList.Add(this);
        if (firstDelay == false) {
            TriggerAndCheck();
        }
    }


    // 触发并检查
    private void TriggerAndCheck() {
        this.times += 1;
        if (this.onTimer != null) {
            this.onTimer(this.times);
        }
        
        if (this.repeat == -1 || this.times < this.repeat){
            this.leftTime = during;
        }
        else {
            Stop();
        }       
    }


    // 取消定时器
    public void Stop() {
        timerList.Remove(this);
    }


    // 定时器步进
    public void Step(float delta) {
        this.leftTime -= delta;
        if (this.leftTime <= 0) {
            TriggerAndCheck();
        }
    }


    // 定时器管理器步进
    public static void Update(float delta) {
        for (int i = 0; i < timerList.Count; i++) {
            timerList[i].Step(delta);
        }
    }

}
