require 'docker'

module ExecCode
  class Container
    def initialize(image, cmd)
      @container = create(image, cmd)
    end

    def run
      begin
        stdout, stderr = @container.tap(&:start).attach(stdout: true, stderr: true)
      rescue
        nil, nil
      ensure
        delete
      end
    end

    private

    def create(image, cmd)
      begin
        Docker::Container.create('Cmd' => [cmd], 'Image' => image)
      rescue
        nil
      end
    end

    def delete
      @container.delete(force: true) unless @container.nil?
    end
  end
end
