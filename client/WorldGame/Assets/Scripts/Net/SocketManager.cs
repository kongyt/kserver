using System;
using System.Net;
using System.Net.Sockets;
using System.Threading;
using UnityEngine;

public class SocketManager : Singleton<SocketManager>
{
    private string ip;
    private int port;

    private bool _isConnected = false;
    public bool IsConnected
    {
        get
        {
            return _isConnected;
        }
    }

    private Socket clientSocket = null;
    private Thread receiveThread = null;
    private byte[] recvBuff = new byte[4096];
    private DataBuffer dataBuff = new DataBuffer(1024);

    private NetMsgCenter msgCenter;

    private void Close()
    {
        if (_isConnected == false)
        {
            return;
        }

        _isConnected = false;

        if (receiveThread != null)
        {
            receiveThread.Abort();
            receiveThread = null;
        }

        if (clientSocket != null && clientSocket.Connected)
        {
            clientSocket.Close();
            clientSocket = null;
        }
    }

    private void ReConnect() {

    }

    public void Connect(string ip, int port) {
        if (IsConnected == true) {
            return;
        }

        msgCenter = MBSingleton.GetInstance("NetMsgCenter") as NetMsgCenter;

        this.ip = ip;
        this.port = port;

        try
        {
            clientSocket = new Socket(AddressFamily.InterNetwork, SocketType.Stream, ProtocolType.Tcp);
            IPAddress addr = IPAddress.Parse(this.ip); // 解析ip地址
            IPEndPoint ipEndpoint = new IPEndPoint(addr, this.port);
            IAsyncResult result = clientSocket.BeginConnect(ipEndpoint, new AsyncCallback(onConnectSuccess), clientSocket);
            bool success = result.AsyncWaitHandle.WaitOne(5000, true);
            if (success == false) {
                onConnectTimeout();
            }
        }
        catch (Exception e) {
            onConnectFail();
        }
    }

    private void onConnectSuccess(IAsyncResult iar) {
        try
        {
            Socket client = (Socket)iar.AsyncState;
            client.EndConnect(iar);

            receiveThread = new Thread(new ThreadStart(onReceiveSocket));
            receiveThread.IsBackground = true;
            receiveThread.Start();
            this._isConnected = true;

            Debug.Log("Socket connected.");
        }
        catch (Exception e) {
            Close();
        }
    }

    private void onConnectFail() {
        Close();
    }

    private void onReceiveSocket() {
        while (true) {
            if (clientSocket.Connected == false) {
                _isConnected = false;
                ReConnect();
                break;
            }

            try{
                int recvLen = clientSocket.Receive(recvBuff);
                if (recvLen > 0)
                {
                    dataBuff.AddBuffer(recvBuff, recvLen);
                    MessageData msgData;
                    while (dataBuff.GetMessage(out msgData)) {
                        msgCenter.PostMessage(msgData);
                    }
                }
            }
            catch (Exception e) {
                Debug.Log("recv msg exception:" + e.StackTrace);
                clientSocket.Disconnect(true);
                clientSocket.Shutdown(SocketShutdown.Both);
                clientSocket.Close();
                break;
            }
        }
    }

    private void onConnectTimeout() {
        Close();
    }


    private void onSendMsg(IAsyncResult asyncSend) {
        try
        {
            Socket client = (Socket)asyncSend.AsyncState;
            client.EndSend(asyncSend);
        }
        catch (Exception e) {
            Debug.Log("send msg exception:" + e.StackTrace);
        }
    }

    public void SendMsg(UInt32 msgId, byte[] msgData) {
        if (clientSocket == null || clientSocket.Connected == false) {
            ReConnect();
            return;
        }

        UInt16 dataLen = (UInt16)(msgData.Length + 4);
        byte[] headLenBuf = BitConverter.GetBytes(IPAddress.HostToNetworkOrder((short)dataLen));
        byte[] msgIdBuf = BitConverter.GetBytes(IPAddress.HostToNetworkOrder((int)msgId));
        byte[] data = new byte[2+dataLen];
        Array.Copy(headLenBuf, 0, data, 0, 2); // 总长度
        Array.Copy(msgIdBuf, 0, data, 2, 4); // 消息类型
        Array.Copy(msgData, 0, data, 6, msgData.Length); // 消息字节流
        clientSocket.BeginSend(data, 0, data.Length, SocketFlags.None, new AsyncCallback(onSendMsg), clientSocket);
    }
}
