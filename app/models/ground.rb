require 'tableless'

class Ground
  include SuperActive::Tableless
  
  attr_accessor :language
  attr_accessor :code
	
	def content
    "package main\r\n\r\nimport \"fmt\"\r\n\r\nfunc main() {\r\n\tfmt.Println(\"Hello, playground\")\r\n}\r\n"
  end
end
