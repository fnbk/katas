BowlingGame = require '../lib/bowlingGame'

describe "Bowling Game", ->

  beforeEach ->
    @rollSpare = ->
      @bowlingGame.roll(4)
      @bowlingGame.roll(6)
    @bowlingGame = new BowlingGame

  it "should have initial score 0", ->
    expect(@bowlingGame.score()).toEqual 0

  it "should have score according to rolled pins", ->
    @bowlingGame.roll(2)
    expect(@bowlingGame.score()).toEqual 2

  it "should add multiple non-strike or spare rolls", ->
    @bowlingGame.roll(1)
    @bowlingGame.roll(8)
    @bowlingGame.roll(7)
    expect(@bowlingGame.score()).toEqual 16

  it "should have 10 points after a spare", ->
    @rollSpare()
    expect(@bowlingGame.score()).toEqual 10

  it "should add bonus points after a spare", ->
    @rollSpare()

    @bowlingGame.roll(3)
    expect(@bowlingGame.score()).toEqual 13 + 3

  it "should not add spare bonus if does not exist", ->
    @bowlingGame.roll(4)
    @bowlingGame.roll(4)
    @bowlingGame.roll(2)
    @bowlingGame.roll(4)
    expect(@bowlingGame.score()).toEqual 14

  it "should not add spare bonus if 10 points are consecutive but not in same frame", ->
    @bowlingGame.roll(4)
    @bowlingGame.roll(3)
    @bowlingGame.roll(7)
    @bowlingGame.roll(1)
    expect(@bowlingGame.score()).toEqual 15
