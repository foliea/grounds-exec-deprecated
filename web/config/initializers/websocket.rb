module WebSocket
  extend self

  def run_endpoint
    ENV.fetch('RUN_ENDPOINT')
  end
end
