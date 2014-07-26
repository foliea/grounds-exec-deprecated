require 'timeout'
require 'docker'

module ExecCode
  class Container
    def initialize(image, cmd)
      begin
        Timeout::timeout(5) do
          @container = create(image, cmd)
        end
      rescue
        raise ExecCode::ContainerCreateError
      end
    end

    def valid?
      !@container.nil?
    end

    def run
      begin
        stdout, stderr = Timeout::timeout(10) do
          @container.tap(&:start).attach(stdout: true, stderr: true)
        end
      rescue
        raise ExecCode::ContainerRunError
      ensure
        delete
      end
    end

    private

    def create(image, cmd)
      Docker::Container.create('Cmd' => [cmd], 'Image' => image)
    end

    def delete
      Timeout::timeout(5) do
        @container.delete(force: true) unless @container.nil?
      end
    end
  end
end
