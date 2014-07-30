require 'timeout'
require 'docker'
require 'exec-code/config'

module ExecCode
  module Service
    extend self

    def available?
      return false unless docker_registry?
      return false unless docker_url?
      return false unless docker_running?
      true
    end

    private

    def docker_registry?
      !ExecCode.docker_registry.empty?
    end

    def docker_url?
      !ExecCode.docker_url.empty?
    end

    def docker_running?
      Docker.url = ExecCode.docker_url
      begin
        Timeout::timeout(5) do
          Docker.version
        end 
        true
      rescue
        false
      end
    end
  end
end
