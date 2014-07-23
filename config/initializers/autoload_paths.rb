Dir[File.join(Rails.root, 'lib', 'containers', '**', '*.rb')].each { |path| require path }
