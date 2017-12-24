using System;
using System.Collections.Generic;
using System.Linq;
using System.Text;

public abstract class FSMState{
    public List<FSMTransition> transtions;

    public abstract void OnEnter();
    public abstract void OnUpdate();
    public abstract void OnExit();

}