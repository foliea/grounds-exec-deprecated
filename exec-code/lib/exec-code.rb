require 'config'
require 'container'
require 'language'
require 'sample'

module ExecCode
  extend self

  Output = Struct.new(:stdout, :stderr)

  def run(language, code, &block)
    cmd = format_input(code)
    begin
      image = "#{docker_registry}/exec-#{language}"
      container = ExecCode::Container.new(image, cmd)
      out, err = container.run(&block)
    rescue
      return nil 
    end
    format_output(out, err)
  end

  private

  def format_input(code)
    code.gsub("\r", "\\r")
        .gsub("\n", '\\n')
  end

  def format_output(out, err)
    Output.new(out.join, err.join)
  end
end
