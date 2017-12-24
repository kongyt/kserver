using UnityEngine;
using System.Collections;

public class TimerManager : MonoBehaviour {

	// Use this for initialization
	void Start () {
	
	}
	
	// Update is called once per frame
	void Update () {
        Timer.Update(Time.deltaTime);
	}

    public void Init() {

    }
}
