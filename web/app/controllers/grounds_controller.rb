class GroundsController < ApplicationController
  include ActionController::Live

  def show
    @ground = GroundDecorator.new(Ground.new(session[:language]), view_context)
  end

  def run
    response.headers["Content-Type"] = "text/event-stream"
    5.times do |n|
      response.stream.write "data: #{n}...\n\n"
      sleep 2
    end
    response.stream.write "data: stop\n\n"
  rescue IOError
    logger.info "Stream closed"
  ensure
    response.stream.close
  end
  
  def change_option
    option, value = params[:option], params[:value]
    session[option] = value if option.present? && value.present?
    render json: { status: :ok }
  end
end