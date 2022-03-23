using System;

namespace State2
{
    class Program
    {
        static void Main(string[] args)
        {

            Fan chain = new Fan();

            while (true)
            {
                getLine();
                chain.pull();
            }


            static String getLine()
            {
                Console.WriteLine("Press Enter ");
                String temp = Console.ReadLine();

                return temp;
            }



        }
    }
}
