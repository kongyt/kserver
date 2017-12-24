using UnityEngine;
using System.Collections.Generic;


// 针对MonoBehaviour的单例管理类
//
// 例子:
// public class MBSingletonTest : MonoBehaviour{
//    public void Test(){
//        Debug.Log("Test");
//    }
// }
//
// 调用
// MBSingletonTest test = MBSingleton.GetInstance("MBSingletonTest") as MBSingletonTest;
// test.Test();
//

public class MBSingleton : MonoBehaviour {

	private static GameObject container = null;
	private static string m_name = "MBSingleton";
	private static SortedDictionary<string, object> singletonMap = new SortedDictionary<string, object>();
	private static bool isDestroying = false;

	public static bool IsDestroying(){
		return isDestroying;
	}

	// 如果一定要在OnDestroy()里调用单例，可以用这个函数判断单例是否存在
	public static bool IsCreatedInstance(string name){
		if (container == null) {
			return false;		
		}
		if (singletonMap != null && singletonMap.ContainsKey (name)) {
			return true;		
		}
		return false;
	}

	public static object GetInstance(string name){
		if (container == null) {
			Debug.Log("Create Singleton Container.");
			container = new GameObject();
			container.name = MBSingleton.m_name;
			container.AddComponent(typeof(MBSingleton));
		}

		if (!singletonMap.ContainsKey (name)) {
			if(System.Type.GetType(name) != null){
				singletonMap.Add(name, container.AddComponent(System.Type.GetType(name)));
			}else{
				Debug.LogWarning("Singleton Type Error! (" + name + ")");
			}
		}
		return singletonMap[name];
	}

	public static void RemoveInstance(string name){
		if (container != null && singletonMap.ContainsKey (name)) {
			UnityEngine.Object.Destroy((UnityEngine.Object)(singletonMap[name]));
			singletonMap.Remove(name);

			Debug.LogWarning("Singleton Remove! (" + name + ")");
		}
	}

	void Awake(){
		Debug.Log ("Awake Singleton.");
		DontDestroyOnLoad (gameObject);
	}

	// Use this for initialization
	void Start () {
		Debug.Log ("Start Singleton.");
	}
	
	// Update is called once per frame
	void Update () {
		
	}

	void OnApplicationQuit(){
		Debug.Log("Destroy Singleton.");
		if (container != null) {
			GameObject.Destroy(container);
			container = null;
			isDestroying = true;
		}

	}
}
