class MiaSpectator

  start: ->
    @createSocket()

  register: (serverMessageReceived) ->
    dgram = require("dgram")
    @client = dgram.createSocket("udp4")

    message = new Buffer("REGISTER;florian")
    @client.send message, 0, message.length, 9000, "localhost"

    @client.on "message", serverMessageReceived

  close: ->
    @client.close()

  createSocket: ->

module.exports = MiaSpectator
