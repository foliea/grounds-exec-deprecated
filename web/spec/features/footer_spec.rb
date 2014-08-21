require 'spec_helper'
require 'capybara/rails'

describe 'footer', type: :feature do
  before(:each) do
    visit('/')
  end

  it 'has a link to about page' do
    path = '/about'
    visit_footer_link(path)
    expect(current_path).to eq(path)
  end

  context 'visit external urls' do
    it 'has a link to contact project developers on github' do
      expect_footer_external_url('https://www.github.com/folieadrien/grounds')
    end

    it 'has a link to github project repository' do
      expect_footer_external_url('https://github.com/folieadrien/grounds/issues/new')
    end
  end

  def visit_footer_link(path)
    within(:css, 'footer') do
      find("a[href=\"#{path}\"]").click
    end
  end
  
  def expect_footer_external_url(url)
    footer = find('footer')
    expect(footer).to have_selector("a[href=\"#{url}\"]")
  end
end