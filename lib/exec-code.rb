module ExecCode

  Output = Struct.new(:stdout, :stderr)

  module Launcher
    extend self

    def run(language, code)
      code = format_input(code)
      container = Container.new("foliea/exec-#{language}", code)
      
      out, err = container.run
      return nil if out.nil? && err.nil?
      
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
end
