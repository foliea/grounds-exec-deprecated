# -*- encoding: utf-8 -*-
require File.expand_path('../lib/exec-code/version', __FILE__)

Gem::Specification.new do |s|
  s.name        = 'exec-code'
  s.version     = ExecCode::VERSION
  s.platform    = Gem::Platform::RUBY
  s.summary     = 'ExecCode'
  s.description = 'ExecCode: Run code in sandboxes environments with docker.'
  s.authors     = ['Adrien Folie']
  s.email       = 'folie.adrien@gmail.com'
  s.license     = 'MIT'
  s.homepage    = 'https://github.com/folieadrien/grounds/exec-code'

  s.add_development_dependency 'bundler', '~> 1.0'
  s.add_development_dependency 'rspec', '~> 2.3'
  s.add_development_dependency 'rake', '~> 10.3', '>= 10.3.2'
  s.add_development_dependency 'pry', '~> 0.10', '>= 0.10.0'

  s.add_dependency 'docker-api'

  s.files         = `git ls-files`.split("\n")
  s.test_files    = `git ls-files -- spec/*`.split("\n") 
  s.require_path = 'lib'
end