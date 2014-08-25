require 'spec_helper'

describe 'option selection in ground editor', type: :feature do
  let(:options) { FactoryGirl.build(:options) }

  it 'saves selected options in session' do
    options.each do |option, code|
      visit(ground_show_path)
      select_option(option, code)
      expect_option_in_session(option, code)
    end
  end

  def expect_option_in_session(option, code)
    session_option = page.get_rack_session_key(option)
    expect(session_option).to eq(code)
  end

  def select_option(option, code)
    find("a[data-#{option}=#{code}]").click
  end
end
