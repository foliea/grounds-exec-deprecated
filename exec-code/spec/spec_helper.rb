require 'docker'

registry = ENV['DOCKER_REGISTRY'] || ''

if registry.empty?
  puts 'Please set DOCKER_REGISTRY first.'
  exit
end

docker_host = ENV['DOCKER_HOST'] || ''
docker_url = ENV['DOCKER_URL'] || ''

if docker_host.empty? && docker_url.empty?
  puts 'Please set DOCKER_HOST or DOCKER_URL first.'
  exit
end

def docker_running?
  begin
    Timeout::timeout(5) do
    end 
    true
  rescue
    false
  end
end

unless docker_running?
  puts 'Please run docker first.'
  exit
end

