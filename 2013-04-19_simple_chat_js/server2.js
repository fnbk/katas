// httpd ... Hypertext Transfer Protocol Daemon -> webserver
var httpd = require('http').createServer(handler);
var io = require('socket.io').listen(httpd);
var fs = require('fs');

var redisPort = 6379;
//var redisHostname = 'Florians-MacBook-Pro.local';
var redisHostname = 'localhost';

var redis = require('redis'),
    RedisStore = require('socket.io/lib/stores/redis'),
    pub    = redis.createClient(redisPort, redisHostname),
    sub    = redis.createClient(redisPort, redisHostname),
    client = redis.createClient(redisPort, redisHostname);

io.set('store', new RedisStore({
    redisPub : pub,
    redisSub : sub,
    redisClient : client
}));

httpd.listen(4002);

function handler(req, res) {
    fs.readFile(__dirname + '/index2.html',
    function(err, data) {
        if (err) {
            res.writeHead(500);
            return res.end('Error loading index.html');
        }

        res.writeHead(200);
        res.end(data);
    });
}

io.sockets.on('connection', function(socket) {
    socket.on('clientMessage', function(content) {
        // echo msg back to sender
        socket.emit('serverMessage', 'You said: ' + content);

        // get the username from current socket connection
        socket.get('username', function(err, username) {
            if (! username) {
                username = socket.id;
            }

            socket.get('room', function(err, room) {
                if (err) { throw err; }
                var broadcast = socket.broadcast;
                var message = content;
                if (room) {
                    broadcast.to(room);
                }
                // send to all connected clients, except the one that send the message
                broadcast.emit('serverMessage', username + ' said: ' + message);
            });
        });
    });

    socket.on('login', function(username) {
        socket.set('username', username, function(err) {
            if (err) { throw err; }
            socket.emit('serverMessage', 'Currently logged in as ' + username);
            socket.broadcast.emit('serverMessage', 'User ' + username + ' logged in');
        });
    });

    socket.on('disconnect', function(username) {
        socket.get('username', function(err, username) {
            if(! username) {
                username = socket.id;
            }
            socket.broadcast.emit('serverMessage', 'User ' + username + ' disconnected');
        });
    });

    socket.on('join', function(room) {
        socket.get('room', function(err, oldRoom) {
            if (err) { throw err; }

            socket.set('room', room, function(err) {
                if (err) { throw err; }
                socket.join(room);
                if (oldRoom) {
                    socket.leave(oldRoom);
                }
                socket.get('username', function(err, username) {
                    if (! username) {
                        username = socket.id;
                    }
                });
                socket.emit('serverMessage', 'You joined room ' + room);
                socket.get('username', function(err, username) {
                    if (! username) {
                        username = socket.id;
                    }
                    socket.broadcast.to(room).emit('serverMessage', 'User ' + username + ' joined this room');
                });
            });
        });
    });

    socket.emit('login');
});
