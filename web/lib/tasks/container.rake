task :run => :environment do
  redis_url = ENV['REDIS_PORT'].gsub('tcp', 'redis')

  port = ENV['RAILS_PORT']

  docker_ip = `netstat -nr | grep '^0\.0\.0\.0' | awk '{print $2}'`.gsub("\n", '')
  docker_url = "http://#{docker_ip}:2375"
  sh "REDIS_URL=\"#{redis_url}\" DOCKER_URL=\"#{docker_url}\" bundle exec rails server -p #{port}"
end

task :test => :environment do
  sh 'bundle exec rspec'
end
