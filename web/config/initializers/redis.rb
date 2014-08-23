$redis = Redis::Namespace.new('grounds', redis: Redis.new(url: ENV['REDIS_URL']))
