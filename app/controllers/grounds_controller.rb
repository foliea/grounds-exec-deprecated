class GroundsController < ApplicationController
  before_action :set_editor_theme, only: :show

  def show
    @ground = GroundDecorator.new(Ground.new(language: 'golang'), view_context)
  end

  def run  
  end
  
  def switch_editor_theme
    session[:editor_theme] = params[:editor_theme]
  end
  
  private
  
  def set_editor_theme
    session[:editor_theme] ||= 'chrome'
	end
end
