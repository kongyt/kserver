using UnityEngine;
using System.Collections;
using System.IO;
using System;

using msg;

public class Client : MonoBehaviour
{

    public string ip = "127.0.0.1";
    public int port = 3563;

    private NetMsgCenter msgCenter;

    // Use this for initialization
    void Start()
    {
        msgCenter =  MBSingleton.GetInstance("NetMsgCenter") as NetMsgCenter;

        msgCenter.addObserver((UInt32)EMsg.S2C_Register_Res_ID, new NetMsgHandle(onRegisterRes));
    }

    private bool flag = false;

    // Update is called once per frame
    void Update()
    {
        if (Input.GetKeyDown(KeyCode.C))
        {
            this.Connect();
        }
        else if (Input.GetKeyDown(KeyCode.S)) {
            flag = true;
            this.SendRegisterReq();
        }
    }

    void Connect(){
        SocketManager.GetInstance().Connect(ip, port);
    }

    void SendRegisterReq() {

        C2S_Register_Req req = new C2S_Register_Req();
        req.userName = "kongyt";
        req.password = "password";


        MemoryStream ms = new MemoryStream();
        ProtoBuf.Serializer.Serialize(ms, req);

        SocketManager.GetInstance().SendMsg((UInt32)EMsg.C2S_Register_Req_ID, ms.ToArray());
    }

    public void onRegisterRes(MessageData msgData) {
        S2C_Register_Res res = ProtoBuf.Serializer.Deserialize<S2C_Register_Res>(new MemoryStream(msgData.msgData));

        Debug.Log("register result:" + res.result);
    }
}
