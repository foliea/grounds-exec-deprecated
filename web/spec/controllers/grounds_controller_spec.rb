require 'spec_helper'

describe GroundsController do

  context "when option doesn't exist" do

    it "doesn't save option in session" do
      put(:switch_option, option: 'language', value: 'unknown')

      expect(session['language']).to be_nil
    end
  end

  context "when option type doesn't exist" do

    it "doesn't save option in session" do
      put(:switch_option, option: 'unknown', value: 'ruby')

      expect(session['unknown']).to be_nil
    end
  end

  context 'when option exists' do
    it "saves option in session" do
      option, value = 'language', 'ruby'

      put(:switch_option, option: option, value: value)

      expect(session[option]).to eq(value)
    end
  end
end
