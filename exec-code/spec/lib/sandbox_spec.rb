require 'spec_helper'
require 'exec-code/sandbox'

describe ExecCode::Sandbox do
  let(:language) { 'ruby' }
  let(:code)     { 'puts "lol"' }
 
  it 'initializes' do
    sandbox = ExecCode::Sandbox.new(language, code)
    expect(sandbox).to be_valid 
  end

  it 'executes code' do
    sandbox = ExecCode::Sandbox.new(language, code)
    enter_block = false
    sandbox.execute do |stream, chunk|
      enter_block = true
    end
    expect(enter_block).to eq true
  end
end
