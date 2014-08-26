require 'spec_helper'

describe 'ground sharing' do
  let(:storage) { $redis }
  let(:ground) { FactoryGirl.build(:ground) }

  context 'when sharing a new ground' do
    before(:each) do
      visit(ground_show_path)
    end

    it 'displays a link to this ground shared url', js: :true do
      share
      expect(shared_url).not_to be_empty
    end
  end

  context 'when accessing a shared ground' do
    before(:each) do
      ground.save
      visit(ground_shared_path(ground))
    end

    it 'ground editor content is equal to ground code' do
      content = find('#ground_editor').text
      expect(content).to eq(ground.code)
    end

    it 'ground editor data language is equal to ground language' do
      expect_data('language', ground.language)
    end

    it 'selected language label is equal to ground language' do
      expect_selected_label('language', ground.language)
    end

    it 'generates the same shared url', js: :true do
      share
      expect(URI(shared_url).path).to eq(ground_shared_path(ground))
    end
  end

  context 'when accessing a non-existent shared ground' do
    it 'raises an error' do
      expect { visit(ground_shared_path(0)) }.to raise_error
    end
  end

  def share
    find('#share').click
  end

  def shared_url
    find('input[name="sharedURL"]').value
  end
end
