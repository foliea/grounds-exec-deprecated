require 'spec_helper'
require 'capybara/rails'

describe 'ground editor', type: :feature do
  before(:each) do
    visit(ground_show_path)
  end

  it 'initialize data language from session' do
    ok = ground_data_exist?('language', 'ruby')
    expect(ok).to be true
  end

  it 'initialize data theme from session' do
    ok = ground_data_exist?('theme', 'monokai')
    expect(ok).to be true
  end

  it 'initialize data indent from session' do
    ok = ground_data_exist?('indent', 'tab')
    expect(ok).to be true
  end

  it 'initialize selected language label from session' do
    expect_selected_label('language', 'ruby')
  end

  it 'initialize selected theme label from session' do
    expect_selected_label('theme', 'monokai')
  end

  it 'initialize selected indent label from session' do
    expect_selected_label('indent', 'tab')
  end

  def ground_data_exist?(option, code)
    select_and_refresh(option, code)

    true if find("#ground_editor[data-#{option}=#{code}]")
  end

  def expect_selected_label(option, code)
    select_and_refresh(option, code)

    label = selected_option_label(option, code)
    expect(label).to eq(option_label(option, code))
  end

  def selected_option_label(option, code)
    find("##{option}-name").text
  end

  def select_and_refresh(option, code)
    find("a[data-#{option}=#{code}]").click
    visit(ground_show_path)
  end

  def option_label(option, code)
    GroundEditor.option(option, code)[:label]
  end
end
