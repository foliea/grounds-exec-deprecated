class GroundsController < ApplicationController
  def show
    @ground = GroundDecorator.new(Ground.new, view_context)
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
    switch_option(:theme, params[:code])
    render json: { status: :ok }
  end
  
  def switch_indent
    switch_option(:indent, params[:code])
    render json: { status: :ok }
  end

  private 
  
  def switch_option(option, code)
    if GroundEditor.has_option?(option, code)
      session[option] = GroundEditor.option(option, params[:code])
    end 
  end
end
