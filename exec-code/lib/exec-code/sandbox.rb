require 'timeout'
require 'docker'
require 'exec-code/config'
require 'exec-code/utils'
require 'exec-code/error'

module ExecCode
  class Sandbox
    def initialize(language, code)
      img = "#{ExecCode.docker_registry}/exec-#{language}"
      cmd = ExecCode::Utils.format_input(code)
      Docker.url = ExecCode.docker_url

      @container = create(img, cmd)
    end

    def valid?
      !@container.nil?
    end

    def execute(&block)
      begin
        if block_given?
          @container.tap(&:start).attach(&block)
        end
      rescue
        nil
      ensure
        delete
      end 
    end

    def interrupt
      begin
        Timeout::timeout(5) do
          @container.kill(signal: 'SIGKILL')
        end
      rescue
        nil
      ensure
        delete
      end
    end

    private

    def create(img, cmd)
      begin
        Timeout::timeout(5) do
          Docker::Container.create('Cmd' => [cmd], 'Image' => img)
        end
      rescue
        nil
      end
    end

    def delete
      begin
        Timeout::timeout(5) do
          @container.delete(force: true)
        end
      rescue
        nil
      end
    end
  end
end
