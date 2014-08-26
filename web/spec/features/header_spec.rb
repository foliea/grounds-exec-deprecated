require 'spec_helper'

describe 'header' do
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
  
  it 'has a link to github project repository' do
    expect_header_external_url('https://github.com/folieadrien/grounds')
  end

  def click_header_link(path)
    within(:css, 'header') do
      find("a[href=\"#{path}\"]").click
    end
  end

  def expect_header_external_url(url)
    header = find('header')
    expect(header).to have_selector("a[href=\"#{url}\"]")
  end
end