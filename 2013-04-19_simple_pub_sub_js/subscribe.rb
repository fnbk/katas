require 'rubygems'
require 'redis'

l = ARGV.length-1
channels = ARGV[0..l]
puts channels

$redis = Redis.new(:timeout => 0)
$redis.subscribe(channels.to_a) do |on|
  on.message do |channel, msg|
    puts "#{channel}: #{msg}"
  end
end
