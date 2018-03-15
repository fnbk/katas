class ApplicationController < ActionController::Base
  protect_from_forgery

  def start
    render :template => 'start'
  end

  def viewer_url
    if Static.viewer_url.nil?
      render :text => "please make a 'create' request first"
    else
      render :text => "<iframe width='480' height='360' src='#{Static.viewer_url}' frameborder='0' allowfullscreen></iframe>"
    end
  end

  def applet_html
    if Static.applet_html.nil?
      render :text => "please make a 'create' request first"
    else
      render :text => Static.applet_html
    end
  end

  def create
    options = {
        'accountid' => params['accountid'],
        'authtoken' => params['authtoken']
    }
    response = Screenleap.make_create_request(options)
    response_body = JSON.parse(response.body)
    Static.screen_share_code = response_body["screenShareCode"]
    Static.viewer_url = response_body["viewerUrl"]
    Static.applet_html = response_body["appletHtml"]
    render :json => response_body
  end

  def stop
    if Static.screen_share_code.nil?
      render :json => {"error" => "please make a 'create' request first"}
    else
      options = {
          'accountid' => params['accountid'],
          'authtoken' => params['authtoken']
      }
      response = Screenleap.make_stop_request(Static.screen_share_code, options)
      render :json => {}
    end
  end

  def status
    if Static.screen_share_code.nil?
      render :json => {"error" => "please make a 'create' request first"}
    else
      options = {
          'accountid' => params['accountid'],
          'authtoken' => params['authtoken']
      }
      response = Screenleap.make_status_request(Static.screen_share_code, options)
      response_body = JSON.parse(response.body)
      render :json => response_body
    end
  end

  def account
    options = {
        'accountid' => params['accountid'],
        'authtoken' => params['authtoken']
    }
    response = Screenleap.make_account_request(options)
    response_body = JSON.parse(response.body)
    render :json => response_body
  end
end
