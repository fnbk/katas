var redis = require('redis');
var client = redis.createClient();

process.argv.slice(2).forEach(function (channel, i) {

    client.subscribe(channel, function () {
        console.log('Subscribing to ' + channel + ' channel');
    });
});

client.on('message', function (channel, msg) {
    console.log("%s: %s", channel, msg);
});
