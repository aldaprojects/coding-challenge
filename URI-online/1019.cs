// Conversión de Tiempo
// https://www.urionlinejudge.com.br/judge/es/problems/view/1019

using System;

class URI {

    static void Main ( string[] args ) {

        int initialSeconds = Int32.Parse( Console.ReadLine() );

        int hours = initialSeconds / 3600;
        int minutes = initialSeconds % 3600 / 60;
        int seconds = initialSeconds % 3600 % 60;

        Console.WriteLine( $"{hours}:{minutes}:{seconds}" );

    }

}
