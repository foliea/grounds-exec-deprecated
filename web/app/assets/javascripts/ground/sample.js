var samples = [
  ['ruby', 'ruby', 'puts "Hello world"'],
  ['golang', 'golang', 'package main\r\n\r\nimport "fmt"\r\n\r\nfunc main() {\r\n\tfmt.Println("Hello world")\r\n}'],
  ['python2', 'python', 'print "Hello World"'],
  ['python3', 'python', 'print("Hello World")'],
  ['c', 'c_cpp', '#include <stdio.h>\r\n\r\nint main()\r\n{\r\n\tprintf("Hello World\\n");\r\n\treturn 0;\r\n}'],
  ['cpp', 'c_cpp', '#include <iostream>\r\n\r\nint main()\r\n{\r\n\tstd::cout << "Hello World\\n";\r\n\treturn 0;\r\n}'],
  ['csharp', 'csharp', 'using System;\r\n\r\nclass Program\r\n{\r\n\tstatic void Main(string[] args)\r\n\t{\r\n\t\tConsole.WriteLine("Hello world");\r\n\t}\r\n}']
]

function GetMode(language) {
  return getForLanguage(language, 1);
}

function GetSample(language) {
  return getForLanguage(language, 2);
}

function getForLanguage(language, key) {
  for (var i = 0; i < samples.length; i++) { 
    if (samples[i][0] === language) {
      return samples[i][key];
    }
  }
  return '';
}

function formatEndpoint(dockerUrl, containerId) {
  var url = dockerUrl.replace('http', 'ws');
  return url + "/v1.14/containers/" + containerId + '/attach/ws?stderr=1&stdout=1&stream=1';
}
