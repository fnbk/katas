var redis = require('redis');
var client = redis.createClient();
var params = { channel: process.argv[2]};

client.on('ready', function() {

    if (params.channel) {
        client.smembers('Channel:' + params.channel, function (err, keys) {
            keys.forEach(function (key) {
                var field = "message";
                client.hget(key, field, function (err, value) {
                    console.log('%s: %s', params.channel, value);
                });
            });
            client.quit();
        });
        return;
    }
    client.quit();
});
