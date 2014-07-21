require 'tableless'

class Ground
  include SuperActive::Tableless
  
  attr_accessor :language
  attr_accessor :code
end
