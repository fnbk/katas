require "screenleap/version"
require 'net/https'
require 'net/http'
require 'uri'

module Screenleap

  def self.make_create_request(options)
    url = "https://api.screenleap.com/v1/screen-shares"
    return make_screenleap_request(:post, url, options)
  end

  def self.make_stop_request(screen_share_code, options)
    url = "https://api.screenleap.com/v1/screen-shares/#{screen_share_code}/stop"
    return make_screenleap_request(:post, url, options)
  end

  def self.make_status_request(screen_share_code, options)
    url = "https://api.screenleap.com/v1/screen-shares/#{screen_share_code}"
    return make_screenleap_request(:get, url, options)
  end

  def self.make_account_request(options)
    url = "https://api.screenleap.com/v1/screen-shares"
    return make_screenleap_request(:get, url, options)
  end

  private

  def self.make_screenleap_request(request_type, url, options)
    uri = URI.parse(url)
    https = Net::HTTP.new(uri.host, uri.port)
    https.use_ssl = true
    request = Net::HTTP::Get.new(uri.path)  unless /get/.match(request_type.to_s).nil?
    request = Net::HTTP::Post.new(uri.path) unless /post/.match(request_type.to_s).nil?
    options.each do |key, value|
      request.add_field(key, value.to_s)
    end
    return https.request(request)
  end
end
