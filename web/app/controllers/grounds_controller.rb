require 'json'

class GroundsController < ApplicationController

  def show
    @ground = GroundDecorator.new(Ground.new(session[:language]), view_context)
  end
  
  def switch_option
    option, value = params[:option], params[:value]
    if option.present? && value.present? && GroundEditor.has_option?(option, value[:code])
      session[option] = value
    end
    render json: { status: :ok }
  end
end
