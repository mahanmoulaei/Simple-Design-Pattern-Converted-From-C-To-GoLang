using System;
using System.Collections.Generic;
using System.Text;

namespace State2
{
    class Fan
    {

        //its just a state wrapper

        private State currentState;

        public Fan()
        {
            currentState = new Off();
        }

        public void setState(State s)
        {
            currentState = s;
        }


        public void pull()
        {
            currentState.pull(this);
        }

    }
}
