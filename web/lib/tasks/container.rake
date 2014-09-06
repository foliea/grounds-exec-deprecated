REDIS_URL    = (ENV['REDIS_1_PORT']  || 'http://127.0.0.1:6379').gsub('tcp', 'redis')
                                                                .gsub('http', 'redis')
WEBSOCKET_URL = ENV['WEBSOCKET_URL'] || 'http://127.0.0.1:5000'
PORT         = ENV['RAILS_PORT']   || 3000

task :run => :environment do
  if production?
    assets_precompile
  end
  sh run 
end

task :test => :environment do
  sh 'bundle exec rspec'
end

def assets_precompile
  sh 'bundle exec rake assets:precompile'
end

def production?
  ENV['RAILS_ENV'] == 'production'
end

def run
  cmd = ''
  cmd << "REDIS_URL=#{REDIS_URL} "
  cmd << "WEBSOCKET_URL=#{WEBSOCKET_URL} "

  cmd << "bundle exec rails server -p #{PORT}"
end
