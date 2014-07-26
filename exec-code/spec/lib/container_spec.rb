require 'spec_helper'
require 'container'

describe ExecCode::Container do
  it 'creates a container' do
    container = ExecCode::Container.new('golang', '42')
    expect(container).to be_valid 
  end
end
