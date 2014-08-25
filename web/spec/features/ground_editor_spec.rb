require 'spec_helper'

describe 'ground editor', type: :feature do
  let(:options) { FactoryGirl.build(:options) }

  before(:each) do
    visit(ground_show_path)
  end
  
  it 'has no JavaScript errors', js: true do
    expect(page).not_to have_errors
  end

  context 'when refreshing ground editor' do
    it 'initialize data options from session' do
      options.each do |option, code|
        select_option(option, code)
        refresh
        expect_data(option, code)
      end
    end

    it 'initialize selected options labels from session' do
      options.each do |option, code|
        select_option(option, code)
        refresh
        expect_selected_label(option, code)
      end
    end
  end
  
  context 'when selecting an option' do
    it 'changes options labels', js: true do
      options.each do |option, code|
        select_option(option, code)
        expect_selected_label(option, code)
      end
    end

    it 'saves selected options in session' do
      options.each do |option, code|
        visit(ground_show_path)
        select_option(option, code)
        expect_option_in_session(option, code)
      end
    end
  end

  def expect_data(option, code)
    data = find("#ground_editor[data-#{option}=#{code}]")
    expect(data).not_to be_nil
  end

  def expect_selected_label(option, code)
    label = selected_option_label(option, code)
    expect(label).to eq(option_label(option, code))
  end
  
  def expect_option_in_session(option, code)
    session_option = page.get_rack_session_key(option)
    expect(session_option).to eq(code)
  end
  
  def refresh
    visit(ground_show_path)
  end
  
  def select_option(option, code)
    find("a[data-#{option}=#{code}]").click
  end

  def selected_option_label(option, code)
    find("##{option}-name").text
  end

  def option_label(option, code)
    GroundEditor.option(option, code)[:label]
  end
end
