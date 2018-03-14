MiaSpectator = require '../lib/miaSpectator'

describe "mia Spectator", ->
  describe "UDP Socket", ->
    it "should create a socket", ->
      miaSpectator = new MiaSpectator()
      spyOn(miaSpectator, 'createSocket')
      miaSpectator.start()
      expect(miaSpectator.createSocket).toHaveBeenCalled()

    it "should send a message with correct parameter on the socket", ->
      miaSpectator = new MiaSpectator()
      client = jasmine.createSpyObj('client', ['send']);
      spyOn(miaSpectator, 'client').andReturn(client)
      miaSpectator.register()
      expect(client.send).toHaveBeenCalled()




  xdescribe "Registration", ->
