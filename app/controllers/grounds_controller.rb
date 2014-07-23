
class GroundsController < ApplicationController
  def show
    @ground = GroundDecorator.new(Ground.new(language: 'golang'), view_context)
  end

  def run
    cmd = ExecCode::Launcher.run(params[:ground][:language], params[:ground][:code])
    respond_to do |format|
      format.js do
        if cmd.nil?
          render json: { stdout: cmd.stdout, stderr: cmd.stderr, status: :ok }
        else
          render json: { status: :service_unavailable }
        end
      end
    end
  end
  
  def switch_editor_theme
    session[:editor_theme] = params[:editor_theme]
  end
end
