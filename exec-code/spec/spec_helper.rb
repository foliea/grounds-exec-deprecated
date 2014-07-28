require 'config'
require 'service'

ExecCode.config do
  docker_registry ENV['DOCKER_REGISTRY']
  docker_url ENV['DOCKER_INSTANCE']
end

unless ExecCode::Service.available?
  puts 'Please verify specs configuration and if docker is running.'
  exit
end

