module ExecCode
  module Sample
    def self.golang
      "package main\r\n\r\nimport \"fmt\"\r\n\r\nfunc main() {\r\n\tfmt.Println(\"Hello world\")\r\n}\r\n"
    end

    def self.ruby
      "puts \"Hello world\"\r\n"
    end
  end
end
