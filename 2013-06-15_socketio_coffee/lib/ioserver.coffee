server = require('http').Server()
socketIo = require('socket.io')

handleMessage = (messageCommand, messageArgs, socket) ->
  if messageCommand == 'ECHO'
    message = {
      'command' : 'ECHO_ECHO',
      'args' : messageArgs
    }
    socket.emit 'message', JSON.stringify(message)
  else
    message = {
      'command' : 'OTHER_COMMAND',
      'args' : ['other', 'command']
    }
    socket.emit 'message', JSON.stringify(message)

handleSocketConnection = (socket) ->
  console.log("connection received: " + socket.id)
  socket.on 'message', (data) ->
    console.log("message received: " + data)
    message = JSON.parse data
    handleMessage message.command, message.args, socket

exports.start = ->
  console.log "starting"
  io = socketIo.listen(server)
  io.set('log level', 1) # 0 error, 1 warn, 2 info, 3 debug
  io.sockets.on 'connection', handleSocketConnection
  server.listen(9000)
exports.stop = ->
  console.log "stopping"
  server.close()
