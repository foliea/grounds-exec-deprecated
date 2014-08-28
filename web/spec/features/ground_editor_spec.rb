require 'spec_helper'

describe 'ground editor' do
  include GroundControls
  include GroundExpectations

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
    
    it 'has code editor cursor on last line', js: :true do
      expect(editor_cursor_on_last_line?).to be true
    end
    
    it 'has no visible link to this ground shared url' do
      expect_shared_url_visibility(false)
    end
  end

  context 'when selecting an option' do
    it 'updates options dropdowns labels', js: :true do
      options.each do |option, code|
        show_dropdown(option)
        select_option(option, code)
        expect_selected_label(option, code)
      end
    end
    
    it 'updates code editor options', js: :true do
      options.each do |option, code|
        show_dropdown(option)
        select_option(option, code)
        expect_editor_option(option, code)
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
end
