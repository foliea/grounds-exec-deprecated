require 'exec-code/service'

ExecCode.config do
  docker_registry ENV['DOCKER_REGISTRY']
  docker_url ENV['DOCKER_INSTANCE']
end

unless ExecCode::Service.available?
  abort('Please verify specs configuration and if docker is running.')
end
