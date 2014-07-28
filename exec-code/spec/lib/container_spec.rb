require 'spec_helper'
require 'container'

describe ExecCode::Container do
  let(:image) { "#{ExecCode.docker_registry}/exec-golang" }

  context 'when image exist' do
    it 'creates a container' do
      container = ExecCode::Container.new(image, '42')
      expect(container).to be_valid 
    end

    it 'runs a container' do
      container = ExecCode::Container.new(image, '42')
      out, _ = container.run
      expect(out).to_not be_nil
    end

    it 'runs a container with a block' do
      container = ExecCode::Container.new(image, '42')
      enter_block = false
      container.run do |stream, chunk|
        enter_block = true
      end
      expect(enter_block).to be_true
    end
  end

  context "when image doesn't exist" do
    it 'raises an error during creation' do
      create = proc { ExecCode::Container.new('unknown/unknown', '42') }
      expect(create).to raise_error(ExecCode::ContainerCreateError) 
    end

    it 'raises an error during run' do

    end
  end

  it 'deletes a container after run' do

  end
end
