function Ground(editor, language, theme, indent, socketEndpoint) {
  this.editor = editor;
  this.language = language;
  this.theme = theme;
  this.indent = indent;
  this.socket = null;
  this.socketEndpoint = socketEndpoint;

  this.initEditor();
  this.setCursor();
  this.setLanguage();
  this.setTheme();
  this.setIndent();
  
  this.bindEvents();
}

Ground.prototype.initEditor = function() {
  this.editor.getSession().setUseWrapMode(true);
}

Ground.prototype.setCursor = function() {
  var lastLine = this.editor.session.getLength();
  this.editor.gotoLine(lastLine);
  this.editor.focus();
};

Ground.prototype.setLanguage = function() {
  this.editor.getSession().setMode("ace/mode/" + this.language.code);
  $("#language-name").text(this.language.label);
};

Ground.prototype.setTheme = function() {
  this.editor.setTheme("ace/theme/" + this.theme.code);
  $("#theme-name").text(this.theme.label);
};

Ground.prototype.setIndent = function() {
  if (this.indent.code == "tab") {
    this.editor.getSession().setUseSoftTabs(false);
    this.editor.getSession().setTabSize(8);
  } else {
    this.editor.getSession().setUseSoftTabs(true);
    this.editor.getSession().setTabSize(this.indent.code);
  }
  $("#indent-name").text(this.indent.label);
};

Ground.prototype.cleanConsole = function() {
   $("#console").find("span").each(function() {
      this.remove();
   });
};

Ground.prototype.runCode = function(code) {
  data = JSON.stringify({ language: this.language.code, code: code});
  this.socket.send(data);
};

Ground.prototype.bindEvents = function() {
  var that = this;
  // Refresh language
  $(".language-link").on('click', function(event, date) {
    that.language = $(event.currentTarget).data('language');
    that.setLanguage();
  });
  // Refresh theme
  $(".theme-link").on('click', function(event, date) {
    that.theme = $(event.currentTarget).data('theme');
    that.setTheme();
  });
  // Refresh indentation
  $(".indent-link").on('click', function(event, date) {
    that.indent = $(event.currentTarget).data('indent');
    that.setIndent();
  });
  // Open socket to web server
  this.socket = new WebSocket(this.socketEndpoint);
  this.socket.onmessage = function(event) {
    if (event.data.length) {
      response = JSON.parse(event.data);
      $("#console").append($('<span class="'+ response.stream +'">').text(response.chunk));
    }
  };
  // Form submit
  $("#run").on('click', function(event) {
    that.cleanConsole();
    var code = that.editor.getValue();
    that.runCode(code);
  });
};
