var samples = [
  ['golang', 'package main\r\n\r\nimport \"fmt\"\r\n\r\nfunc main() {\r\n\tfmt.Println(\"Hello world\")\r\n}\r\n'],
  ['ruby', 'puts \"Hello world\"\r\n']
]

function GetSample(language) {
  for (i = 0; i < samples.length; i++) { 
    if (samples[i][0] === language) {
      return samples[i][1];
    }    
  }
  return '';
}
