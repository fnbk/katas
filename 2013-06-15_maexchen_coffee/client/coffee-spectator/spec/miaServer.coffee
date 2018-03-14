class MiaServer

  # see: http://goo.gl/zG6km

  start: (readyCallback) ->
    @server = require('child_process').fork(__dirname + '/../../../server/server.js')
    setTimeout readyCallback, 500 # wait 500ms for the server to start

  stop: (finishedCallback) ->
    @server.kill()
    setTimeout finishedCallback, 500 # wait 500ms for the server to stop

  connected: ->
    @server.connected

module.exports = MiaServer
