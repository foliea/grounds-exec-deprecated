module ExecCode
  module Utils
    extend self

    def format_input(input)
      input.gsub("\r", "\\r")
          .gsub("\n", '\\n')
    end
  end
end
