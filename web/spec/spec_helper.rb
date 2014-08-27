ENV["RAILS_ENV"] ||= 'test'
require File.expand_path("../../config/environment", __FILE__)
require 'rspec/rails'
require 'capybara/rails'
require 'capybara/poltergeist'
require 'rack_session_access/capybara'

Dir[Rails.root.join("spec/support/**/*.rb")].each { |f| require f }

# Checks for pending migrations before tests are run.
# If you are not using ActiveRecord, you can remove this line.
# ActiveRecord::Migration.maintain_test_schema!

$redis = MockRedis.new

RSpec.configure do |config|
  # If you're not using ActiveRecord, or you'd prefer not to run each of your
  # examples within a transaction, remove the following line or assign false
  # instead of true.
  config.use_transactional_fixtures = false

  # Use FactoryGirl instead of fixtures
  config.include(FactoryGirl::Syntax::Methods)

  config.infer_spec_type_from_file_location!

  config.include(GroundExpectations)
  config.include(GroundControls)
end

Capybara.register_driver :poltergeist do |app|
    options = {
      js_errors: true,
    }
    Capybara::Poltergeist::Driver.new(app, options)
end

Capybara.javascript_driver = :poltergeist
