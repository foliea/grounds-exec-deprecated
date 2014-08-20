module WebSocket
  extend self

  def run_endpoint
    ENV['RUN_ENDPOINT']
  end
end
