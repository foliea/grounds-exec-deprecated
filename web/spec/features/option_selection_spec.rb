require 'spec_helper'
require 'capybara/rails'

describe 'option selection in ground editor', type: :feature do
  before(:each) do
    visit('/')
  end

  it 'saves selected language in session' do
    value = 'golang'
    language = switch_option('language', value)
    expect(language).to eq(value)
  end

  it 'saves selected theme in session' do
    value = 'monokai'
    theme = switch_option('theme', value)
    expect(theme).to eq(value)
  end

  it 'saves selected indent in session' do
    value = 'tab'
    indent = switch_option('indent', value)
    expect(indent).to eq(value)
  end

  def switch_option(option, value)
    find("a[data-#{option}=#{value}]").click
    page.get_rack_session_key(option)
  end
end
