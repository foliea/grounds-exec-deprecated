require 'json'

class GroundsController < ApplicationController

  def show
    @ground = GroundDecorator.new(Ground.new(session[:language]), view_context)
  end
  
  def switch_option
    option, value = params[:option], params[:value]
    session[option] = value if option.present? && value.present?
		# verify with ground editor if option is valid
	 	# GroundEditor.has_option?(option, value)
    render json: { status: :ok }
  end
end
