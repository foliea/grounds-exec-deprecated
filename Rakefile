#!/usr/bin/env rake

DOCKER_REGISTRY = 'grounds'
DOCKER_BUILD_SERVER = 'docker build -t grounds-server .'
DOCKER_BUILD_WEB = 'docker build -t grounds-web web'

namespace :test do
  task :server do
    sh DOCKER_BUILD_SERVER
    sh 'docker run --rm grounds-server hack/test.sh'
  end

  task :web do
    sh DOCKER_BUILD_WEB
    sh 'docker run --rm grounds-web bundle exec rspec'
  end
end

namespace :run do
  task :server do
    sh DOCKER_BUILD_SERVER
    sh "docker run -d -p 8080:8080 grounds-server hack/run.sh '-d -r #{DOCKER_REGISTRY}'"
  end

  task :web => :server do
    sh DOCKER_BUILD_WEB
    sh 'docker run -d -p 3000:3000 -e RUN_ENDPOINT=192.168.59.103:8080/ws grounds-web rails s -p 3000'
  end 
end

namespace :images do
  desc 'Build docker images for a given registry'
  task :build, [:registry] do |_, args|
    registry = args[:registry] || DOCKER_REGISTRY
    images do |file, path|
      sh "docker build -t #{registry}/#{file} #{path}"
    end
  end

  desc 'Push docker images to a given registry'
  task :push, [:registry] do |_, args|
    registry = args[:registry] || DOCKER_REGISTRY
    images do |file, _|
      sh "docker push #{registry}/#{file}"
    end
  end

  desc 'Pull docker images from a given registry'
  task :pull, [:registry] do |_, args|
    registry = args[:registry] || DOCKER_REGISTRY
    images do |file, _|
      sh "docker pull #{registry}/#{file}"
    end
  end
end

def images
  return unless block_given?

  dockerfiles = File.join("#{File.dirname(__FILE__)}", 'dockerfiles')

  Dir.entries(dockerfiles).select { |f| f != '.' && f != '..' }.each do |file|
    path = File.join(dockerfiles, file)
    if File.directory?(path)
      yield(file, path)
    end
  end
end
