var samples = [
  ['ruby', 'ruby', 'puts "Hello world"'],
  ['golang', 'golang', 'package main\r\n\r\nimport "fmt"\r\n\r\nfunc main() {\r\n\tfmt.Println("Hello world")\r\n}'],
  ['python2', 'python', 'print "Hello, World"'],
  ['python3', 'python', 'print("Hello, World")'],
  ['c', 'c_cpp', '#include <stdio.h>\r\n\r\nint main()\r\n{\r\n\tprintf("Hello World\\n");\r\n\treturn 0;\r\n}'],
  ['cpp', 'c_cpp', '#include <iostream>\r\n\r\nint main()\r\n{\r\n\tstd::cout << "Hello World\\n";\r\n\treturn 0;\r\n}']
]

function GetTheme(language) {
  for (i = 0; i < samples.length; i++) { 
    if (samples[i][0] === language) {
      return samples[i][1];
    }    
  }
  return '';
}

function GetSample(language) {
  for (i = 0; i < samples.length; i++) { 
    if (samples[i][0] === language) {
      return samples[i][2];
    }    
  }
  return '';
}

