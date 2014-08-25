require 'spec_helper'

describe 'ground sharing' do
  let(:storage) { $redis }
  let(:ground) { FactoryGirl.build(:ground) }

  context 'when accessing a shared ground' do
    before(:each) do
      ground.save
      visit(ground_shared_path(ground))
    end

    it 'ground_editor content is equal to ground code' do
      content = find('#ground_editor').text
      expect(content).to eq(ground.code)
    end

    it 'ground_editor data language is equal to ground language' do
      ok = find("#ground_editor[data-language=#{ground.language}]")
      expect(ok).not_to be_nil
    end
  end

  context 'when accessing a non-existent shared ground' do
    it 'raises an error' do
      expect { visit(ground_shared_path(0)) }.to raise_error
    end
  end
end
