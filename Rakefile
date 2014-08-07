#!/usr/bin/env rake

task :build do
  sh './hack/server.sh ./bin'
end

default_registry = 'grounds'

namespace :images do
  desc 'Build docker images for a given registry (default: grounds)'
  task :build, [:registry] do |_, args|
    registry = args[:registry] || default_registry
    images do |file, path|
      sh "docker build -t #{registry}/#{file} #{path}"
    end
  end

  desc 'Push docker images to a given registry (default: grounds)'
  task :push, [:registry] do |_, args|
    registry = args[:registry] || default_registry
    images do |file, _|
      sh "docker push #{registry}/#{file}"
    end
  end

  desc 'Pull docker images from a given registry (default: grounds)'
  task :pull, [:registry] do |_, args|
    registry = args[:registry] || default_registry
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
