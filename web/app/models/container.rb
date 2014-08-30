require 'docker'

class Container
  def initialize(internal)
    @internal = internal
  end

  def id
    @internal.id
  end

  def url
    @internal.connection.url
  end

  def start
    @internal.start
  end

  def stop
    @internal.stop
  end

  def status
    @internal.wait['StatusCode']
  end

  def remove
    @internal.delete(force: true)
  end

  def self.find_by_id(id)
    new(Docker::Container.get(id))
  end

  def self.create(language, code)
    image, cmd = format_image(language), format_cmd(code)
    new(Docker::Container.create('Cmd' => [cmd], 'Image' => image ))
  end

  private

  def self.format_image(language)
    "grounds/exec-#{language}"
  end

  def self.format_cmd(code)
    code.gsub("\\", "\\\\\\")
        .gsub("\r", "\\r")
        .gsub("\n", "\\n")
        .gsub("\t", "\\t")
  end
end
