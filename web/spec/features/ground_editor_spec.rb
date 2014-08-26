require 'spec_helper'

describe 'ground editor' do
  let(:options) { FactoryGirl.build(:options) }

  before(:each) do
    visit(ground_show_path)
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
    it 'changes options labels', js: :true do
      options.each do |option, code|
        show_dropdown(option)
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

  def refresh
    visit(ground_show_path)
  end

  def show_dropdown(option)
    find("a[data-dropdown=#{option.pluralize(2)}]").click
  end

  def select_option(option, code)
    find("a[data-#{option}=#{code}]").click
  end
end
