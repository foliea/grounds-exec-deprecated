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

  def switch_editor_theme
    session[:editor_theme] = params[:theme] if params[:theme].present?
    render json: { theme: session[:editor_theme], status: :ok }
  end
end
