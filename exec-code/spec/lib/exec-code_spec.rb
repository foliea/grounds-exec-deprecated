require 'spec_helper'
require 'exec-code'

describe ExecCode::Launcher do
  it 'gets a compilation error from container' do
    output = ExecCode::Launcher.run('golang', '42')
    expect(output.stderr).to_not be_empty
  end

  it 'gets hello world from container' do
    code = ExecCode::Sample.from('golang')
    output = ExecCode::Launcher.run('golang', code)
    expect(output.stdout).to_not be_empty
  end

  context 'when block given to run' do
    it 'enter the block' do
      code = ExecCode::Sample.from('golang')
      enter_block = false
      ExecCode::Launcher.run('golang', code) do |stream, chunk|
        enter_block = true
      end
      expect(enter_block).to be_true
    end

    it 'returns an output' do
      code = ExecCode::Sample.from('golang')
      output = ExecCode::Launcher.run('golang', code) {}
      expect(output).to_not be_nil
    end
  end
end
