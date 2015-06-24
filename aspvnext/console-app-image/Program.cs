using System;
using System.Threading;

public class Program
{
    public static void Main()
    {
    	while(true){

    		Console.WriteLine("Hello World : {0}", DateTime.UtcNow.ToLongTimeString() );
    		Thread.Sleep(5000);

    	}       
    }
}