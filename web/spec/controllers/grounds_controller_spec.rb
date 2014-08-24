require 'spec_helper'

describe GroundsController do

  context "when option doesn't exist" do

    it "doesn't save option in session" do
      put(:switch_option, option: 'language', code: 'unknown')

      expect(session['language']).to be_nil
    end
  end

  context "when option type doesn't exist" do

    it "doesn't save option in session" do
      put(:switch_option, option: 'unknown', code: 'ruby')

      expect(session['unknown']).to be_nil
    end
  end

  context 'when option exists' do
    it "saves option in session" do
      option, code = 'language', 'ruby'

      put(:switch_option, option: option, code: code)

      expect(session[option]).to eq(code)
    end
  end
end
