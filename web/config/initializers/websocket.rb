module WebSocket
  extend self

  def run_endpoint
    ENV['WS_RUN_ENDPOINT']
  end
end
