task :run => :environment do
  redis_url    = ENV['REDIS_PORT'].gsub('tcp', 'redis')
  run_endpoint = ENV['RUN_ENDPOINT']
  port         = ENV['RAILS_PORT']

  run(redis_url, run_endpoint, port)
end

task :test => :environment do
  sh 'bundle exec rspec'
end

def run(redis_url, run_endpoint, port)
  if ENV['RAILS_PORT'] == 'production'
    sh 'bundle exec rake:assets precompile'
  end
  run_command = "REDIS_URL='#{redis_url}'"
  run_command << " RUN_ENDPOINT='#{run_endpoint}'"
  run_command << " bundle exec rails server -p #{port}"
  sh run_command
end
