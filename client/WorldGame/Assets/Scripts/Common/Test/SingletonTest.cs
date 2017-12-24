using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public class SingletonTest : Singleton<SingletonTest> {

	public void Test(){
		Debug.Log("Singleton run ok.");
	}
}
