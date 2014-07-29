class GroundsController < ApplicationController
  include Tubesock::Hijack

  def show
    @ground = GroundDecorator.new(Ground.new(session[:language]), view_context)
  end
  
  def run
    hijack do |tubesock|
      tubesock.onmessage do |data|
        tubesock.send_data("You're' trying to execute: #{data}")
        5.times do
          tubesock.send_data("Faking execution of code.")
          sleep 2
        end
      end
      tubesock.onclose do
        # Kill container if still alive
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
