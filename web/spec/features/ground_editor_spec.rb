require 'spec_helper'
require 'capybara/rails'

describe 'ground editor', type: :feature do
  let(:options) { FactoryGirl.build(:options) }

  context 'when visiting ground editor' do
    before(:each) do
      visit(ground_show_path)
    end

    it 'initialize data options from session' do
      options.each do |option, code|
        select_and_refresh(option, code)
        expect(ground_data_exist?(option, code)).to be true
      end
    end

    it 'initialize selected options labels from session' do
      options.each do |option, code|
        select_and_refresh(option, code)
        expect_selected_label(option, code)
      end
    end
  end

  def ground_data_exist?(option, code)
    true if find("#ground_editor[data-#{option}=#{code}]")
  end

  def expect_selected_label(option, code)
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
