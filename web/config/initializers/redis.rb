$redis = Redis::Namespace.new('grounds', redis: Redis.new(url: ENV.fetch('REDIS_URL')))
