Rover = require '../lib/marsRover'

describe 'Mars rover', ->

  moveWithTable = (table_test, move) ->
    table_test.forEach (tt) ->
      rover = new Rover(tt.in)
      move(rover)
      expect(rover.position).toEqual {x: tt.out.x, y: tt.out.y}
      expect(rover.direction).toEqual tt.out.d

  beforeEach ->
    @rover = new Rover(x: 0, y: 0, d: 'N')

  it "should have the correct start position", ->
    expect(@rover.position).toEqual {x: 0, y: 0}
    expect(@rover.direction).toEqual 'N'

  it "should move to the correct forward position with given direction", ->
    table_test = [
      {in: {x: 0, y: 0, d: 'N'}, out: {x: 0, y: 1, d: 'N'}},
      {in: {x: 0, y: 0, d: 'E'}, out: {x: 1, y: 0, d: 'E'}},
      {in: {x: 0, y: 0, d: 'W'}, out: {x: -1, y: 0, d: 'W'}},
      {in: {x: 0, y: 0, d: 'S'}, out: {x: 0, y: -1, d: 'S'}},
    ]
    moveWithTable table_test, (rover) -> rover.moveForward()

  it "should move to the correct backward position with given direction", ->
    table_test = [
      {in: {x: 0, y: 0, d: 'N'}, out: {x: 0, y: -1, d: 'N'}},
      {in: {x: 0, y: 0, d: 'E'}, out: {x: -1, y: 0, d: 'E'}},
      {in: {x: 0, y: 0, d: 'W'}, out: {x: 1, y: 0, d: 'W'}},
      {in: {x: 0, y: 0, d: 'S'}, out: {x: 0, y: 1, d: 'S'}},
    ]
    moveWithTable table_test, (rover) -> rover.moveBackward()
