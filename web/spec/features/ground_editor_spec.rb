require 'spec_helper'

describe 'ground editor' do
  include GroundControls
  include GroundExpectations

  let(:options) { TestOptionsTable }

  before(:each) do
    visit(ground_show_path)
  end

  it 'has no visible link to a shared url' do
    expect_shared_url_visibility(false)
  end

  context 'when first visit to a new ground' do
    it 'initialize data options from default option' do
      options.each do |option, _|
        expect_data(option, default_option_code(option))
      end
    end
    
    it 'initialize selected options labels from default option' do
      options.each do |option, _|
        expect_selected_label(option, default_option_code(option))
      end
    end
    
    it 'initialize code editor options from default option', js: :true do
      options.each do |option, _|
        expect_editor_option(option, default_option_code(option))
      end
    end
  end

  context 'when selecting an option and refreshing ground editor' do
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

    it 'initialize code editor options from session', js: :true do
      options.each do |option, code|
        show_dropdown(option)
        select_option(option, code)
        refresh
        expect_editor_option(option, code)
      end
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
        expect(session(option)).to eq(code)
      end
    end
    
    it 'closes properly the dropdown associated', js: :true do
      options.each do |option, code|
        show_dropdown(option)
        select_option(option, code)
        expect(dropdown_closed?(option)).to be true
      end
    end
  end
end
