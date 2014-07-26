# -*- encoding: utf-8 -*-
require File.expand_path('../lib/exec-code/version', __FILE__)

Gem::Specification.new do |s|
  s.name        = 'exec-code'
  s.version     = ExecCode::VERSION
  s.platform    = Gem::Platform::RUBY
  s.summary     = 'ExecCode'
  s.description = 'Run code in sandboxes environments'
  s.authors     = ['Adrien Folie']
  s.email       = 'folie.adrien@gmail.com'
  s.license     = 'MIT'
  s.homepage    = 'https://github.com/folieadrien/exec-code'

  s.add_development_dependency 'bundler', '~> 1.0'
  s.add_development_dependency 'rspec', '~> 2.3'
  s.add_development_dependency 'pry', '~> 0.10', '>= 0.10.0'
  s.add_development_dependency 'rake', '~> 10.3', '>= 10.3.2'

  s.add_dependency 'docker-api'

  s.files = [
    'lib/container.rb',
    'lib/error.rb',
    'lib/exec-code.rb',
    'lib/image.rb',
    'lib/sample.rb'
  ]
  s.require_path = 'lib'
end
