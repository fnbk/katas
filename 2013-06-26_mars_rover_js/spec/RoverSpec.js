var Rover = require('../lib/Rover.js').Rover;

describe("Rover should", function() {
  // Prepare for tests
  beforeEach(function () {
	this.rover = new Rover({x: 0, y: 0, d: 'N'});
  });
  
  afterEach(function () {
	this.rover = null;
  });
  
  it("have initial position", function () {
      this.r1 = new Rover({x: 0, y: 0, d: 'N'});
      expect(this.r1.position).toEqual({x: 0, y: 0});
      this.r2 = new Rover({x: 1, y: 1, d: 'N'});
      expect(this.r2.position).toEqual({x: 1, y: 1});
  });
  
  it("have initial direction", function () {
      this.r1 = new Rover({x: 0, y: 0, d: 'N'});
      expect(this.r1.direction).toEqual('N');
      this.r2 = new Rover({x: 0, y: 0, d: 'E'});
      expect(this.r2.direction).toEqual('E');
  });
  
  it("be able to move forwards", function () {
	expect(this.rover.position).toEqual({x: 0, y: 0});
	this.rover.forward();
	expect(this.rover.position).toEqual({x: 0, y: -1});
  });
  
  it("be able to move backwards", function () {
     expect(this.rover.position).toEqual({x: 0, y: 0}); 
     this.rover.backward();
     expect(this.rover.position).toEqual({x: 0, y: 1});
  });
  
  it("accept single character commands", function () {
	spyOn(this.rover, "forward");
	expect(this.rover.forward).not.toHaveBeenCalled();
	this.rover.parse("f");
	expect(this.rover.forward).toHaveBeenCalled();
  });
});