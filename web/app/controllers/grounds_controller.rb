class GroundsController < ApplicationController
  include ActionController::Live

  def show
    @ground = GroundDecorator.new(Ground.new(session[:language]), view_context)
  end
 
  def change_option
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
