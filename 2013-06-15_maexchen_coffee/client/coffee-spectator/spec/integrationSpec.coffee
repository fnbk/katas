MiaServer = require './miaServer'
MiaSpectator = require '../lib/miaSpectator'

xdescribe "Coffee Spectator", ->

  it "should start the mia server", (done) ->
    miaServer = new MiaServer()
    miaServer.start ->
      miaServer.stop ->
        done()

  it "should register at mia server", (done) ->
    miaServer = new MiaServer()
    miaServer.start ->
      miaSpectator = new MiaSpectator()
      serverMessageReceivedCallback = (msg, rinfo) ->
        response = "client got: #{msg} from #{rinfo.address}:#{rinfo.port}"
        expect(typeof response).toEqual 'string'
        miaSpectator.close()
        miaServer.stop ->
          done()
      miaSpectator.register serverMessageReceivedCallback
