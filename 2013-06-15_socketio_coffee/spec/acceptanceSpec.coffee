io = require('socket.io-client')
ioserver = require('../lib/ioserver')

describe "smoke test", ->
  it "should pass", ->
    expect(true).toBe true

describe "Client", ->
  app = null
  beforeEach (done) ->
    ioserver.start()
    done()

  afterEach (done) ->
    ioserver.stop()
    done()

  it "should be able to connect", (done) ->
    client = io.connect('http://localhost:9000')
    client.on 'connect', (data) ->
      expect(true).toBe(true)
      done()

  it "should receive ECHO_ECHO message", (done) ->
    message = {
      'command' : 'ECHO',
      'args'    : ['hello', 'echo']
    }
    client = io.connect('http://localhost:9000')
    client.emit('message', JSON.stringify(message))
    client.on 'message', (data) ->
      expect(JSON.parse(data).command).toEqual 'ECHO_ECHO'
      done()
