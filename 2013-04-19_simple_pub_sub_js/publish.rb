require 'rubygems'
require 'redis'

channel = ARGV[0]
message = ARGV[1]
#puts "channel: #{channel}"
#puts "message: #{message}"

$redis = Redis.new
$redis.publish channel, message

rand_key = (Random.rand(12000) * Random.rand(12000)).to_i().to_s(16).sub('.', '')
$redis.hset(rand_key, "message", message)
$redis.sadd('Channel:' + channel, rand_key)

