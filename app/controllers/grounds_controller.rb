class GroundsController < ApplicationController
  def show
    @ground = GroundDecorator.new(Ground.new('golang'), view_context)
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
  end
  
  def switch_indent
    switch_option(:indent, params[:code])
  end
  
  private 
  
  def switch_option(option, code)
    options = GroundEditor.send(option.to_s.pluralize(2))
    if options.has_key?(code)
      session[option] = { code: code, label: options[code] }
    end
    render json: { status: :ok }
  end
end
