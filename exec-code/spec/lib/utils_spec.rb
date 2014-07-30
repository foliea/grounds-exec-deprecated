require 'exec-code/utils'

describe ExecCode::Utils do
  it 'formats input string with tabs and carriage return' do
    input = "puts \"Hello world\"\r\n\t"
    desired = "puts \"Hello world\"\\r\\n\\t"
    expect(ExecCode::Utils.format_input(input)).to eq(desired)
  end
  
  it 'formats input string with \\#{char} to escape' do
    input = "puts \"Hello world\\n\"\r\n"
    desired = "puts \"Hello world\\\\n\"\\r\\n"
    expect(ExecCode::Utils.format_input(input)).to eq(desired)
  end
end
