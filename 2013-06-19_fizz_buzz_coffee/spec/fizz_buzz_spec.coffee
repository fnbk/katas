FizzBuzz = require '../lib/fizz_buzz'

describe "FizzBuzz", ->
  beforeEach ->
    @fizzBuzz = new FizzBuzz
    @array = @fizzBuzz.getArray(100)

  it "should return 1 as the first number", -> 
    expect(@array[0]).toEqual 1
      
  it "should return 2 as the second number", ->
    expect(@array[1]).toEqual 2

  it "should return 'Fizz' for the third number", ->
    expect(@array[2]).toEqual "Fizz"

  it "should return 'Fizz' for every third number (except for numbers also divisible by 5)", ->
    for i in [1..100]
      if i % 3 == 0 and i % 5 != 0 and not @fizzBuzz.number_has_five_in_it(i)
        expect(@array[i-1]).toEqual "Fizz"

  it "should return 'Buzz' for the fifth number", ->
    expect(@array[4]).toEqual "Buzz"
  
  it "should return 'FizzBuzz' for the fifteenth number", ->
    expect(@array[14]).toEqual "FizzBuzz"
    
  it "should return 'Fizz' for 13", ->
    expect(@array[12]).toEqual "Fizz"
    
  it "should return 'FizzBuzz' for 30", ->
    expect(@array[29]).toEqual "FizzBuzz"
    
  it "should return 'FizzBuzz' for 51", ->
    expect(@array[50]).toEqual "FizzBuzz"