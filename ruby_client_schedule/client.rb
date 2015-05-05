require 'sidekiq'

Sidekiq.configure_client do |config|
  config.redis = { :url => "redis://localhost:6379", :namespace => 'goworkers-sidekiq' }
end

class HelloWorker
  include Sidekiq::Worker

  def perform(name = nil)
    if name
      puts "Hello #{name}"
    else
      puts "Hello World!"
    end
  end
end

HelloWorker.perform_in 5
