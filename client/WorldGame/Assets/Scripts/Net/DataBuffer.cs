using System.IO;
using System;
using UnityEngine;

public struct MessageData
{
    public UInt16 msgLen;
    public UInt32 msgId;
    public byte[] msgData;
}

// 自动大小的数据缓存
public class DataBuffer
{
    private int minBuffLen;
    private byte[] buff;
    private int curBuffPostion;

    public DataBuffer(int minBuffLen = 1024)
    {
        if (minBuffLen <= 0)
        {
            this.minBuffLen = 1024;
        }
        else
        {
            this.minBuffLen = minBuffLen;
        }
        buff = new byte[this.minBuffLen];
    }

    public void AddBuffer(byte[] data, int dataLen)
    {
        if (dataLen > buff.Length - curBuffPostion)
        {// 超过当前缓存
            byte[] tmpBuf = new byte[curBuffPostion + dataLen];
            Array.Copy(buff, 0, tmpBuf, 0, curBuffPostion);
            Array.Copy(data, 0, tmpBuf, curBuffPostion, dataLen);
            tmpBuf = null;
        }
        else
        {
            Array.Copy(data, 0, buff, curBuffPostion, dataLen);
        }
        curBuffPostion += dataLen; // 修改当前数据标记
    }

    public bool GetMessage(out MessageData msgData)
    {
        bool rs = false;
        msgData = new MessageData();

        if (curBuffPostion >= 2)
        {
            UInt16 dataLen = BitConverter.ToUInt16(buff, 0);
            int totalDataLen = 2 + dataLen;
            if (totalDataLen <= curBuffPostion)
            {                
                msgData.msgLen = (ushort)(dataLen - 4);
                msgData.msgId = BitConverter.ToUInt32(buff, 2);
                msgData.msgData = new byte[msgData.msgLen];
                Array.Copy(buff, 6, msgData.msgData, 0, msgData.msgLen);

                int leftDataLen = curBuffPostion - totalDataLen;
                curBuffPostion = leftDataLen;
                byte[] tmpBuff = new byte[leftDataLen < minBuffLen ? minBuffLen : leftDataLen];

                Array.Copy(buff, dataLen, tmpBuff, 0, leftDataLen);
                buff = tmpBuff;
                rs = true;
            }
        }
        return rs;
    }
}