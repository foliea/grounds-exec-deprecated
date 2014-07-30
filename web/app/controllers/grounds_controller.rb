class GroundsController < ApplicationController
  include Tubesock::Hijack

  def show
    @ground = GroundDecorator.new(Ground.new(session[:language]), view_context)
  end
  
  def run
    hijack do |sock|
      sandbox = nil
      sock.onmessage do |data|
        h = JSON.parse(data)
        # sandbox.interrupt unless sandbox.nil?
        # works with closure ?
        # -> try without sandbox = nil on top
        # -> not working: try with instance variable
        sandbox = ExecCode::Sandbox.new(h['language'], h['code'])
        sandbox.execute do |stream, chunk|
          sock.send_data(chunk)
        end
      end
      sock.onclose do
      # sandbox.interrupt unless sandbox.nil?
      end
    end
  end

  def switch_option
    option, value = params[:option], params[:value]
    if option.present? && value.present?
      session[option] = value
      if option == 'language'
        custom = ExecCode::Sample.from(value[:code])
      end
    end
    render json: { status: :ok, custom: custom }
  end
end
