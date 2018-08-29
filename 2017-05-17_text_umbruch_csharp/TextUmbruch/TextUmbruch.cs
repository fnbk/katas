namespace TextUmbruch
{
    using System;

    public class TextUmbruch
    {
        public static string Umbrechen(string text, int maxZeilenLaenge)
        {
            return text;
        }

        public static string[] WoerterBilden(string text)
        {
            return text.Split(new [] {' ', '\n'}, StringSplitOptions.RemoveEmptyEntries);
        }
    }
}













































//var splitted = text.Split(' ');
//return splitted.Where(s => !string.IsNullOrWhiteSpace(s)).ToArray();
