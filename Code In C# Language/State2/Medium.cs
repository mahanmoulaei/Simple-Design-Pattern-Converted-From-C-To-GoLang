using System;
using System.Collections.Generic;
using System.Text;

namespace State2
{
    class Medium : State
    {
        public void pull(Fan wrapper)
        {
            wrapper.setState(new High());
            Console.WriteLine("turning High speed");
        }
    }
}
