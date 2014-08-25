task :run => :environment do
  redis_url = ENV['REDIS_PORT'].gsub('tcp', 'redis')
  run_endpoint = ENV['WEBSOCKET_PORT'].gsub('tcp', 'ws') + '/run'

  port = ENV['RAILS_PORT']

  sh "REDIS_URL=\"#{redis_url}\" RUN_ENDPOINT=\"#{run_endpoint}\" bundle exec rails server -p #{port}"
end

task :test => :environment do
  sh 'xvfb-run bundle exec rspec'
end