#!/usr/bin/env rake

$LOAD_PATH.unshift File.expand_path('./lib', __FILE__)

require 'rspec/core/rake_task'
require 'images'
require 'docker'

RSpec::Core::RakeTask.new(:spec)

task :default => :spec

namespace :images do
  task :build do
 
    ExecCode::Images::all.each do |name|
      puts "Building image: #{name}..."
      Docker::Image.build_from_dir("./dockerfiles/#{name}/",
                                   t: "foliea/#{name}:latest") do |chunk| 
        puts chunk
      end
      puts "Built image #{name} with success!"
    end
  end

  task :push do
    #authenticate!
    ExecCode::Images::all.each do |name|
      puts "Pushing: #{name} to docker hub..."
      puts "Pushed #{name} to docker hub with success!"
    end
  end
end

def authenticate!
  puts 'Trying to authenticate to docker hub...'
  begin 
    Docker.authenticate!(username: ENV['DOCKER_USERNAME'], 
                         password: ENV['DOCKER_PASSWORD'], 
                         email:    ENV['DOCKER_EMAIL'])
  rescue
    puts 'Failed to authenticate to docker hub...'
    exit
  end
  puts 'Authenticated to docker hub with success!'
end
