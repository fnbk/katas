using Microsoft.VisualStudio.TestTools.UnitTesting;

namespace TextUmbruch
{
    [TestClass]
    public class TextUmbruchTest
    {
        [TestMethod]
        public void TestUmbruchSimple()
        {
            const string input = @"a\nb\nc";
            var output = TextUmbruch.Umbrechen(input, 1);
            Assert.AreEqual(input, output);
        }

        [TestMethod]
        public void TestWoerterBilden_Simple()
        {
            const string input = "a b c";
            var expected = new[] { "a", "b", "c"};
            var actual = TextUmbruch.WoerterBilden(input);
            CollectionAssert.AreEqual(expected, actual);
        }

        [TestMethod]
        public void TestWoerterBilden_DoubleSpace()
        {
            const string input = "a  b c  ";
            var expected = new[] { "a", "b", "c" };
            var actual = TextUmbruch.WoerterBilden(input);
            CollectionAssert.AreEqual(expected, actual);
        }

        [TestMethod]
        public void TestWoerterBilden_Commas()
        {
            const string input = "a ,b c,";
            var expected = new [] { "a", ",b", "c," };
            var actual = TextUmbruch.WoerterBilden(input);
            CollectionAssert.AreEqual(expected, actual);
        }

        [TestMethod]
        public void TestWoerterBilden_CommasAsWords()
        {
            const string input = "a , b c,";
            var expected = new[] { "a", "," ,"b", "c," };
            var actual = TextUmbruch.WoerterBilden(input);
            CollectionAssert.AreEqual(expected, actual);
        }

        [TestMethod]
        public void TestWoerterBilden_Newlines()
        {
            const string input = "\na\n\nb\n";
            var expected = new[] { "a", "b" };
            var actual = TextUmbruch.WoerterBilden(input);
            CollectionAssert.AreEqual(expected, actual);
        }
    }
}
