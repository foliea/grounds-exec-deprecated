module ExecCode
  module Sample
    extend self

    def from(language)
      send(language) if respond_to?(language, true)
    end

    private

    def golang
      "package main\r\n\r\nimport \"fmt\"\r\n\r\nfunc main() {\r\n\tfmt.Println(\"Hello world\")\r\n}\r\n"
    end

    def ruby
      "puts \"Hello world\"\r\n"
    end
  end
end
