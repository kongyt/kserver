using UnityEngine;
using System.Collections;
using System;
using System.Collections.Generic;

public class ObjectPoll<T>{

    private readonly Stack<T> stack = new Stack<T>();
    private readonly Func<T> actionOnNew;
}
