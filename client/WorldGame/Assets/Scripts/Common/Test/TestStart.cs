using UnityEngine;
using System.Collections;

public class TestStart : MonoBehaviour {

    MBSingletonTest singletonTest;
    TimerManager timerManager;

    void Awake() {
        singletonTest = MBSingleton.GetInstance("MBSingletonTest") as MBSingletonTest;
        timerManager = MBSingleton.GetInstance("TimerManager") as TimerManager;
        timerManager.Init();
    }

	// Use this for initialization
	void Start () {

		SingletonTest.GetInstance ().Test ();
        singletonTest.Test2();

        new Timer(1, 10, true, new TimerEventHandler(this.OnTimer));
	}
	
	// Update is called once per frame
	void Update () {
       
	}

    void OnTimer(int times) {
        Debug.Log("定时器触发(" + times + ")");
    }
}
