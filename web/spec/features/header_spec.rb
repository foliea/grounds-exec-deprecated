require 'spec_helper'
require 'capybara/rails'

describe 'header', type: :feature do
  before(:each) do
    visit('/')
  end

  it 'has a link to site root' do
    path = '/'
    visit_header_link(path)
    expect(current_path).to eq(path)
  end

  it 'has a link to about page' do
    path = '/about'
    visit_header_link(path)
    expect(current_path).to eq(path)
  end

  def visit_header_link(path)
    within(:css, 'header') do
      find("a[href=\"#{path}\"]").click
    end
  end
end