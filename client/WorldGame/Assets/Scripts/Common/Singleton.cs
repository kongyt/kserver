using UnityEngine;
using System.Collections;

// 普通单例模板类
// 
// 例子
// public class SingletonTest : Singleton<SingletonTest>{
//     public void Test(){
//         Debug.Log("Test");   
//     }
// }
// 
// 调用
// SingletonTest.GetInstance().Test();
//


public class Singleton<T> where T:new() {

	private static T instance = default(T);
	public static T GetInstance(){
		if (instance == null) {
			instance = new T ();
		}
		return instance;
	}
}
