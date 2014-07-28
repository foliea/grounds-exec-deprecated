class GroundsController < ApplicationController
  include ActionController::Live

  def show
    @ground = GroundDecorator.new(Ground.new(session[:language]), view_context)
  end

  def run
  end
  
  def change_option
    option, value = params[:option], params[:value]
    session[option] = value if option.present? && value.present?
    render json: { status: :ok }
  end
end