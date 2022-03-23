using System;
using System.Collections.Generic;
using System.Text;

namespace State2
{
    class Off : State
    {
        public void pull(Fan wrapper)
        {
            wrapper.setState(new Low());
            Console.WriteLine("turning Low speed");
        }
    }
}
