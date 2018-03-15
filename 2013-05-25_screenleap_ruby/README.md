# Screenleap [![Build Status](https://travis-ci.org/fnbk/screenleap.png)](https://travis-ci.org/fnbk/screenleap) [![Code Climate](https://codeclimate.com/github/fnbk/screenleap.png)](https://codeclimate.com/github/fnbk/screenleap)

Simple screenleap API wrapper. See: http://www.screenleap.com/api

## Installation

Add this line to your application's Gemfile:

    gem 'screenleap'

And then execute:

    $ bundle

Or install it yourself as:

    $ gem install screenleap

## Test

To run the tests you need to export accountid and authtoken environment variables, then bundle and then run the tests:

    export SCREENLEAP_ACCOUNTID=abcdef
    export SCREENLEAP_AUTHTOKEN=opqrstuvw
    bundle install
    rake

vcr saves all http requests in spec/fixtures/vcr_cassets/. So after running the tests a second time the execution should be significantly faster, because webmock uses stored responses.




## Usage

This gem is a simple wrapper for the given api requests. To find out more about the options go to http://www.screenleap.com/api/reference

### create request

Example: 

```ruby
options = {
  'accountid'       => "your account id",
  'authtoken'       => "your auth token",
  'isSecure'        => false,
  'showPauseButton' => false
}
response = Screenleap.make_create_request(options)
response_body = JSON.parse(response.body)

puts response_body["screenShareCode"]
puts response_body["viewerUrl"]
puts response_body["appletHtml"]
puts response_body["errorMessage"]
```

### stop request

Example: 

```ruby
options = {
    'accountid' => "your account id",
    'authtoken' => "your auth token"
}
response = Screenleap.make_stop_request("screen share code", options)

puts response #=> ""
```

### status request

Example: 

```ruby
options = {
    'accountid' => "your account id",
    'authtoken' => "your auth token"
}
response = Screenleap.make_status_request("screen share code", options)
response_body = JSON.parse(response.body)

puts response_body["centsPerMinute"]
puts response_body["isSecure"]
puts response_body["durationInMinutes"]
```

### account request

Example: 

```ruby
options = {
  'accountid' => "your account id",
  'authtoken' => "your auth token"
}
response = Screenleap.make_account_request(options)
response_body = JSON.parse(response.body)

puts response_body #=> array
```


## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request
