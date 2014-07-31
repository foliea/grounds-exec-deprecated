require 'json'

class GroundsController < ApplicationController

  def show
    @ground = GroundDecorator.new(Ground.new(session[:language]), view_context)
  end
  
  def switch_option
    option, value = params[:option], params[:value]
    if option.present? && value.present?
      session[option] = value
      if option == 'language'
        #custom = ExecCode::Sample.from(value[:code])
      end
    end
    render json: { status: :ok, custom: nil }
  end
end
