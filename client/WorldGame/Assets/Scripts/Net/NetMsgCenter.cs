using UnityEngine;
using System.Collections;
using System.Collections.Generic;
using System;

public delegate void NetMsgHandle(MessageData msgData);

public class NetMsgCenter : MonoBehaviour
{

    private Dictionary<UInt32, NetMsgHandle> netMsgEventList = new Dictionary<uint, NetMsgHandle>();
    public Queue<MessageData> messageDataQueue = new Queue<MessageData>();

    // 添加网络事件观察者
    public void addObserver(UInt32 msgId, NetMsgHandle handle) {
        if (netMsgEventList.ContainsKey(msgId))
        {
            netMsgEventList[msgId] += handle;
        }
        else {
            netMsgEventList.Add(msgId, handle);
        }
    }

    // 删除网络事件观察者
    public void removeObserver(UInt32 msgId, NetMsgHandle handle) {
        if (netMsgEventList.ContainsKey(msgId)) {
            netMsgEventList[msgId] -= handle;
        }
        if (netMsgEventList[msgId] == null) {
            netMsgEventList.Remove(msgId);
        }
    }

    // Use this for initialization
    void Start()
    {

    }

    // Update is called once per frame
    void Update()
    {
        while (messageDataQueue.Count > 0) {
            lock (messageDataQueue) {
                MessageData msgData = messageDataQueue.Dequeue();
                if (netMsgEventList.ContainsKey(msgData.msgId)) {
                    netMsgEventList[msgData.msgId](msgData);
                }
            }
        }
    }

    public void PostMessage(MessageData msgData) {
        lock (messageDataQueue) {
            messageDataQueue.Enqueue(msgData);
        }
    }
}
