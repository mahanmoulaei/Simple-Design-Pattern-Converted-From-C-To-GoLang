using System;
using System.Collections.Generic;
using System.Text;

namespace State2
{
    interface State
    {

        void pull(Fan wrapper);
    }
}
