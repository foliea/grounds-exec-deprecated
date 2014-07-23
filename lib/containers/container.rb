require 'docker'

module Container
  extend self
 
  Output = Struct.new(:stdout, :stderr)

  def exec(language, code)
    container = create(language, code)
    stdout, stderr = container.tap(&:start).attach(stdout: true, stderr: true)
    container.delete(force: true)
    Output.new(stdout.join, stderr.join)
  end

  private

  def create(language, code)
    Docker::Container.create('Cmd' => [format_code(code)],
                             'Image' => "foliea/exec-#{language}:latest")
  end

  def format_code(code)
    code.gsub("\r", "\\r")
        .gsub("\n", '\\n')
  end
end
