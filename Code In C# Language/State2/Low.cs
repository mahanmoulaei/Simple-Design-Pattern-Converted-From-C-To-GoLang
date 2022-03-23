using System;
using System.Collections.Generic;
using System.Text;

namespace State2
{
    class Low : State
    {
        public void pull(Fan wrapper)
        {
            wrapper.setState(new Medium());
            Console.WriteLine("turning Medium speed");
        }
    }
}
