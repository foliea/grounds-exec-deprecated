class Ground
  attr_reader :language
  attr_reader :code

  def initialize(language = nil)
    @language = language || GroundEditor.default_option(:language)
    @code = "package main\r\n\r\nimport \"fmt\"\r\n\r\nfunc main() {\r\n\tfmt.Println(\"Hello, playground\")\r\n}\r\n"
  end

  def valid?
    GroundEditor.has_option?(:language, @language[:code])
  end
 end
