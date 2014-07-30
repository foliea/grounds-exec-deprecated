module ExecCode
  module Utils
    extend self

    def format_input(input)
      input.gsub("\\", "\\\\\\")
           .gsub("\r", "\\r")
           .gsub("\n", "\\n")
           .gsub("\t", "\\t")
    end
  end
end
