var redis = require('redis');
var client = redis.createClient();
var params = { channel: process.argv[2], message: process.argv[3]};

client.on('ready', function() {

    if (params.channel && params.message) {
        var randKey = "Messages:" + (Math.random() * Math.random()).toString(16).replace('.', '');
        client.hset(randKey, "message", params.message);
        client.sadd('Channel:' + params.channel, randKey);
        client.publish(params.channel, params.message);
    }

    client.quit();
});
