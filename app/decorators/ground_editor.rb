require 'yaml'

module GroundEditor
  extend self
 
  @config = YAML.load_file("#{Rails.root}/config/editor.yml")
  @config.each do |key, value|
    self.send(:define_method, key) { value }
  end
end