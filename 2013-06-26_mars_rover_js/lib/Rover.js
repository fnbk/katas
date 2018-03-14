function Rover (coordinates) {
  this.position = {
	x: coordinates.x,
	y: coordinates.y
  };
  
  this.direction = coordinates.d;
  
  this.forward = function () {
      this.position.y -= 1;
  };
  
  this.backward = function () {
	this.position.y += 1;
  };
  
  this.parse = function(direction) {
    this.forward();
  }
}

exports.Rover = Rover;
