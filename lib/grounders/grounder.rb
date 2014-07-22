require 'docker'

class Grounder
  def initialize
    Docker.url = ENV['DOCKER_URL']
    @registry = 'foliea'
  end

  def run(language, code)
    container = create(language, code)
    stdout, stderr = container.tap(&:start).attach(stdout: true, stderr: true)
    container.delete(force: true)
    [stdout, stderr]
  end

  private

  def create(language, code)
    Docker::Container.create('Cmd' => [format_code(code)], 'Image' => "#{@registry}/#{language}:latest")
  end

  def format_code(code)
    code.gsub("\r", "\\r")
        .gsub("\n", '\\n')
  end
end
