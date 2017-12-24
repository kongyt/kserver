using UnityEngine;
using System.Collections;
using System.Collections.Generic;

public abstract class FSM{

    protected FSMState activeState;

    public void Start(FSMState state) {
        activeState = state;
        activeState.OnEnter();
    }


    public void Update() {

        if (activeState != null) {
            activeState.OnUpdate();

            for (int i = 0; i < activeState.transtions.Count; i++) {
                FSMTransition trans = activeState.transtions[i];
                if (trans.IsValid()) {
                    activeState.OnExit();

                    activeState = trans.GetNextState();
                    trans.OnTransition();

                    activeState.OnEnter();
                    break;
                }
            }

        }
    }

}
