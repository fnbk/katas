DIRECTIONS =
  N:
    moveForward: (position) -> {x: position.x, y: position.y + 1}
    moveBackward: (position) -> {x: position.x, y: position.y - 1}
  E:
    moveForward: (position) -> {x: position.x + 1, y: position.y}
    moveBackward: (position) -> {x: position.x - 1, y: position.y}
  S:
    moveForward: (position) -> {x: position.x, y: position.y - 1}
    moveBackward: (position) -> {x: position.x, y: position.y + 1}
  W:
    moveForward: (position) -> {x: position.x - 1, y: position.y}
    moveBackward: (position) -> {x: position.x + 1, y: position.y}

class Rover
  constructor: (positionWithDirection) ->
    {x, y, d} = positionWithDirection
    @position = {x, y}
    @direction = d

  moveForward: ->
    @position = DIRECTIONS[@direction].moveForward(@position)

  moveBackward: ->
    @position = DIRECTIONS[@direction].moveBackward(@position)


module.exports = Rover