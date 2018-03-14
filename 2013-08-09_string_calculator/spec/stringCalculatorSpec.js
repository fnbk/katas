var String = require('../src/stringCalculator.js');

describe("String", function() {
  it("should have a calc method", function() {
    expect("".calc).toBeTruthy();
  });

  describe("#calc", function() {
    it("should return 0 if the string is empty", function() {
      expect("".calc()).toBe(0);
    });

    it("should return a number from the string", function() {
      expect("3".calc()).toBe(3);
    });

    it("should return the sum of three comma separated numbers", function() {
      expect("1,2,3".calc()).toBe(6);
    });

    it("should work with line breaks", function() {
      expect("1\n,\n\n2,3".calc()).toBe(6);
    });

    it("should work with other delimiters, introduced by a start line with //", function() {
      expect("//;\n1;2;4".calc()).toBe(7);
    });
  });
});
