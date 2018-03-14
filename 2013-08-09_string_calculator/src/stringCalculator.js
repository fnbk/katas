
String.prototype.calc = function() {
  if(this == "") {
    return 0;
  } else {
    if(this.charAt(0) == "/" && this.charAt(1) == "/") {
      var sum = 0;
      var delimiter = this.charAt(2);
      var startPositionOfNumbers = this.indexOf("\n");
      var numbers = this.slice(startPositionOfNumbers+1);
      var arrayOfStrings = numbers.split(delimiter);
      for (var i=0; i<arrayOfStrings.length; i++) {
        sum = sum + parseInt(arrayOfStrings[i]);
      }
      return sum;
    } else {
      var sum = 0;
      var arrayOfStrings = this.split(",");
      for (var i=0; i<arrayOfStrings.length; i++) {
        sum = sum + parseInt(arrayOfStrings[i]);
      }
      return sum;
    }
  }
}

module.exports = String;
