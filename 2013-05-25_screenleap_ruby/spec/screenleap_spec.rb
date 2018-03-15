require 'spec_helper'

# http://www.screenleap.com/api/reference

describe "Screenleap" do
  describe "smoke" do
    it "should not smoke in this test" do
      true.should be true
    end
  end

  describe "#make_create_request" do
    # POST https://api.screenleap.com/v1/screen-shares

    it "should respond to make_create_request with one parameter" do
      Screenleap.should respond_to(:make_create_request).with(1).argument
    end

    it "should make a post request to screenleap", :vcr do
      Screenleap.make_create_request({})
      WebMock.should have_requested(:post, "https://api.screenleap.com/v1/screen-shares")
    end

    it "should make a request with given option values", :vcr do
      options = {
          'accountid'        => ENV['SCREENLEAP_ACCOUNTID'],
          'authtoken'        => ENV['SCREENLEAP_AUTHTOKEN'],
          'isSecure'         => false,
          'showStopButton'   => true,
          'showScreenToggle' => true,
          'showPauseButton'  => false
      }
      Screenleap.make_create_request(options)
      WebMock.should have_requested(:post, "https://api.screenleap.com/v1/screen-shares").with(:headers => options)
    end

    it "should return a subclass of Net::HTTPResponse", :vcr do
      Screenleap.make_create_request({}).should be_kind_of Net::HTTPResponse
    end
  end

  describe "#make_stop_request" do
    # POST https://api.screenleap.com/v1/screen-shares/{screenShareCode}/stop

    it "should respond to make_stop_request with two parameters" do
      Screenleap.should respond_to(:make_stop_request).with(2).arguments
    end

    it "should make the appropriate post request to screenleap", :vcr do
      screen_share_code = "abc"
      Screenleap.make_stop_request(screen_share_code, {})
      WebMock.should have_requested(:post, "https://api.screenleap.com/v1/screen-shares/#{screen_share_code}/stop")
    end

    it "should make a request with given option values", :vcr do
      options = {
          'accountid'        => ENV['SCREENLEAP_ACCOUNTID'],
          'authtoken'        => ENV['SCREENLEAP_AUTHTOKEN']
      }
      screen_share_code = "abc"
      Screenleap.make_stop_request(screen_share_code, options)
      WebMock.should have_requested(:post, "https://api.screenleap.com/v1/screen-shares/#{screen_share_code}/stop").with(:headers => options)
    end

    it "should return a subclass of Net::HTTPResponse", :vcr do
      Screenleap.make_stop_request(nil, {}).should be_kind_of Net::HTTPResponse
    end
  end

  describe "#make_status_request" do
    # GET https://api.screenleap.com/v1/screen-shares/{screenShareCode}

    it "should respond to make_status_request with two parameters" do
      Screenleap.should respond_to(:make_status_request).with(2).arguments
    end

    it "should make the appropriate get request to screenleap", :vcr do
      screen_share_code = "abc"
      Screenleap.make_status_request(screen_share_code, {})
      WebMock.should have_requested(:get, "https://api.screenleap.com/v1/screen-shares/#{screen_share_code}")
    end

    it "should make a request with given option values", :vcr do
      options = {
          'accountid'        => ENV['SCREENLEAP_ACCOUNTID'],
          'authtoken'        => ENV['SCREENLEAP_AUTHTOKEN']
      }
      screen_share_code = "abc"
      Screenleap.make_status_request(screen_share_code, options)
      WebMock.should have_requested(:get, "https://api.screenleap.com/v1/screen-shares/#{screen_share_code}").with(:headers => options)
    end

    it "should return a subclass of Net::HTTPResponse", :vcr do
      Screenleap.make_status_request(nil, {}).should be_kind_of Net::HTTPResponse
    end
  end

  describe "#make_account_request" do
    # GET https://api.screenleap.com/v1/screen-shares

    it "should respond to make_account_request with one parameter" do
      Screenleap.should respond_to(:make_account_request).with(1).argument
    end

    it "should make the appropriate get request to screenleap", :vcr do
      Screenleap.make_account_request({})
      WebMock.should have_requested(:get, "https://api.screenleap.com/v1/screen-shares")
    end

    it "should make a request with given option values", :vcr do
      options = {
          'accountid'        => ENV['SCREENLEAP_ACCOUNTID'],
          'authtoken'        => ENV['SCREENLEAP_AUTHTOKEN'],
          'endedAfter'       => "1348358582702"
      }
      Screenleap.make_account_request(options)
      WebMock.should have_requested(:get, "https://api.screenleap.com/v1/screen-shares").with(:headers => options)
    end

    it "should return a subclass of Net::HTTPResponse", :vcr do
      Screenleap.make_account_request({}).should be_kind_of Net::HTTPResponse
    end
  end
end
