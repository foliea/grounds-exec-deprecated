require 'spec_helper'
require 'exec-code'

describe ExecCode::Launcher do
  it 'gets a compilation error from container' do
    output = ExecCode::Launcher.run('golang', '42')
    expect(output.stderr).to_not be_empty
  end

  it 'gets hello world from container' do
    code = "package main\r\n\r\nimport \"fmt\"\r\n\r\nfunc main() {\r\n\tfmt.Println(\"Hello, playground\")\r\n}\r\n"
    output = ExecCode::Launcher.run('golang', code)
    expect(output.stdout).to_not be_empty
  end
end
