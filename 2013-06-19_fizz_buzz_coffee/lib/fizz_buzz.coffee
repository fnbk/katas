class FizzBuzz
  getArray: ->
    @array = []
    for i in [1..100]
      response = ""
      if i%3 == 0 or @number_has_three_in_it(i)
        response += "Fizz"
      if i%5 == 0 or @number_has_five_in_it(i)
        response += "Buzz"
      unless response
        response = i
      @array.push response
    return @array

  number_has_three_in_it: (number) ->
    if "#{number}".indexOf("3") == -1
      return false
    else
      return true

   number_has_five_in_it: (number) ->
    if "#{number}".indexOf("5") == -1
      return false
    else
      return true

module.exports = FizzBuzz