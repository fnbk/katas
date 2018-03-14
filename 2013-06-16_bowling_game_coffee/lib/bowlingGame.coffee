class BowlingGame
  constructor: ->
    @points = 0
    @rolls = []
    @firstRollInFrame = true

  score: ->
    @points

  roll: (pins) ->
    @points += pins

    if @sumLastTwoRolls() == 10 and @firstRollInFrame
      @points += pins

    @rolls.push pins
    @firstRollInFrame = not @firstRollInFrame

  sumLastTwoRolls: ->
    [z..., secondLast, last] = @rolls
    secondLast + last

module.exports = BowlingGame