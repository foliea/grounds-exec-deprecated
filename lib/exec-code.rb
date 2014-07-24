require 'container'

# FIX: Hardcoded repository
# FIX: Error handling

module ExecCode

  Output = Struct.new(:stdout, :stderr)

  module Launcher
    extend self

    def run(language, code)
      code = format_input(code)
      begin
        container = ExecCode::Container.new("foliea/exec-#{language}", code)
        out, err = container.run
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
end
