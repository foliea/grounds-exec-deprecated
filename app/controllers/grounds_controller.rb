class GroundsController < ApplicationController
  respond_to :html, only: :show

  def show
    @ground = GroundDecorator.new(Ground.new('golang'), view_context)
    respond_with(@ground)
  end

  def run
    cmd = ExecCode::Launcher.run(params[:ground][:language], params[:ground][:code])
    if cmd.nil?
      render json: { status: :service_unavailable }
    else
      render json: { stdout: cmd.stdout, stderr: cmd.stderr, status: :ok }
    end
  end

  def switch_theme
    theme = Theme.get(params[:code])
    session[:editor_theme] = theme if theme.present?
    render json: { status: :ok }
  end
  
  def switch_language
  end
end
