require 'spec_helper'

describe 'ground sharing' do
  let(:storage) { $redis }
  let(:ground) { FactoryGirl.build(:ground) }

  context 'after sharing a new ground' do
    before(:each) do
      visit(ground_show_path)
      share
    end

    it 'has a link to this ground shared url', js: :true do
      expect(shared_url).not_to be_empty
    end

    it 'has a visible link to this ground shared url', js: :true do
      link = find('#sharedURL', visible: true)
      expect(link).not_to be_nil
    end

    context 'when selecting another language' do
      it 'has no visible link to this ground shared url', js: :true do
        show_dropdown('language')
        select_option('language', 'golang')

        link = find('#sharedURL', visible: false)
        expect(link).not_to be_nil
      end
    end

    context 'when typing inside the code editor' do
      it 'has no visible link to this ground shared url', js: :true do
        type_inside_editor

        link = find('#sharedURL', visible: false)
        expect(link).not_to be_nil
      end
    end
  end

  # add test sharedURL is visible:false at start
  # add test sharedURL is visible:false when selecting another language, or typing into editor

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
end
