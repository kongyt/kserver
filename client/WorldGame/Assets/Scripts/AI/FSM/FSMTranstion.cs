using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

public abstract class FSMTransition{
    public abstract bool IsValid();
    public abstract FSMState GetNextState();
    public abstract void OnTransition();
}