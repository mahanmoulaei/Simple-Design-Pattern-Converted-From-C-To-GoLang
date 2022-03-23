using System;
using System.Collections.Generic;
using System.Text;

namespace State2
{
    class High : State
    {
        public void pull(Fan wrapper)
        {
            wrapper.setState(new Off());
            Console.WriteLine("turning Off");
        }
    }
}
