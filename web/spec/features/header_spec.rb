require 'spec_helper'

describe 'header', type: :feature do
  before(:each) do
    visit('/')
  end

  it 'has a link to site root' do
    path = '/'
    click_header_link(path)
    expect(current_path).to eq(path)
  end

  it 'has a link to about page' do
    path = '/about'
    click_header_link(path)
    expect(current_path).to eq(path)
  end

  def click_header_link(path)
    within(:css, 'header') do
      find("a[href=\"#{path}\"]").click
    end
  end
end