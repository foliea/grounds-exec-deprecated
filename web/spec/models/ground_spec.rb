require 'spec_helper'
require 'mock_redis'

describe Ground do
  before(:each) do
    $redis = MockRedis.new
  end

  let(:storage) { $redis }

  it 'is convertible to an hash' do
    ground = FactoryGirl.build(:ground)
    expected = {'language' => ground.language, 'code' => ground.code}
    expect(ground.to_h).to eq(expected)
  end

  it 'generates a key' do
    ground = FactoryGirl.build(:ground)
    expected = 'e5b71bb74c1854f3f264c5332836179e860f14651e8878e6ffc29780596bb221'
    expect(ground.generate_key).to eq(expected)
  end

  context 'when ground is saved' do
    it 'has an id equal to generated key' do
      ground = FactoryGirl.build(:ground)
      ground.save

      key = ground.generate_key
      expect(ground.id).to eq(key)
    end

    it 'is persisted' do
      ground = FactoryGirl.build(:ground)
      ground.save
      expect(ground).to be_persisted
    end

    it 'is exists in storage' do
      ground = FactoryGirl.build(:ground)
      ground.save

      expect(storage.exists(ground.id)).to be true
    end

    it 'can be retrieve from storage' do
      ground = FactoryGirl.build(:ground)
      ground.save

      expected = Ground.from_storage(ground.id)
      expect(ground.to_h).to eq(expected.to_h)
    end

    it 'is destroyable' do
      ground = FactoryGirl.build(:ground)
      ground.save
      ground.destroy
      expect(ground).not_to be_persisted
    end
  end

  context 'when ground is not saved' do
    it 'has no id' do
      ground = FactoryGirl.build(:ground)
      expect(ground.id).to be_nil
    end
  end
end
