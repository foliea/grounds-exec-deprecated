redis = if ENV["RAILS_ENV"] == 'test'
          Redis.new
        else
          Redis.new(url: ENV.fetch('REDIS_URL'))
        end

$redis = Redis::Namespace.new('grounds', redis: redis)
