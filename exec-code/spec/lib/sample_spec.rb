require 'sample'

describe ExecCode::Sample do
  context 'when sample exist' do
    it 'returns the sample' do
      expect(ExecCode::Sample.from('golang')).to_not be_nil
    end
  end

  context "when sample doesn't exist" do
    it 'returns nil' do
      expect(ExecCode::Sample.from('not_exist')).to be_nil
    end
  end
end
