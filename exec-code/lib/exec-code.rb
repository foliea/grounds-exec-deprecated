require 'container'
require 'sample'

# FIX: Hardcoded registry
# FIX: Language not supported / registry empty / image missing

module ExecCode

  Output = Struct.new(:stdout, :stderr)

  module Launcher
    extend self

    def service_available?(language)
      ExecCode::Language.suppored?(language) 
    end

    def run(language, code, &block)
      code = format_input(code)
      begin
        container = ExecCode::Container.new("foliea/exec-#{language}", code)
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
end
