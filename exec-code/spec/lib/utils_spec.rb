require 'exec-code/utils'

describe ExecCode::Sample do
  it 'formats input string' do
    input = "puts \"Hello world\"\r\n"
    desired = "puts \"Hello world\"\\r\\n"
    expect(ExecCode::Utils.format_input(input)).to eq(desired)
  end
end
