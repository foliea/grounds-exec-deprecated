require 'spec_helper'

describe 'ground sharing' do
  include GroundControls
  include GroundExpectations
  
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
      expect_shared_url_visibility(true)
    end

    context 'when selecting another language' do
      it 'has no visible link to this ground shared url', js: :true do
        show_dropdown('language')
        select_option('language', 'golang')
        expect_shared_url_visibility(false)
      end
    end

    context 'when typing inside the code editor' do
      it 'has no visible link to this ground shared url', js: :true do
        type_inside_editor
        expect_shared_url_visibility(false)
      end
    end
  end

  context 'when visiting a shared ground' do
    before(:each) do
      ground.save
      visit(ground_shared_path(ground))
    end

    it 'has code editor content equal to ground code', js: :true do
      expect(editor_content).to eq(ground.code)
    end

    it 'has data equal to shared ground language' do
      expect_data('language', ground.language)
    end

    it 'has selected language label equal to shared ground language' do
      expect_selected_label('language', ground.language)
    end

    it 'generates the same shared url', js: :true do
      share
      expect(URI(shared_url).path).to eq(ground_shared_path(ground))
    end
  end

  context 'when visiting a non-existent shared ground' do
    it 'raises an error' do
      expect { visit(ground_shared_path(0)) }.to raise_error
    end
  end
end
