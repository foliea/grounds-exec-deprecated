require 'spec_helper'
require 'capybara/rails'

describe 'option selection in ground editor', type: :feature do
  before(:each) do
    visit(ground_show_path)
  end

  it 'saves selected language in session' do
    expect_option_in_session('language', 'ruby')
  end

  it 'saves selected theme in session' do
    expect_option_in_session('theme', 'monokai')
  end

  it 'saves selected indent in session' do
    expect_option_in_session('indent', 'tab')
  end

  def expect_option_in_session(option, code)
    select_option(option, code)

    session_option = page.get_rack_session_key(option)
    expect(session_option).to eq(code)
  end

  def select_option(option, code)
    find("a[data-#{option}=#{code}]").click
  end
end
