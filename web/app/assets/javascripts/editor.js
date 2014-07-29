function Ground(editor, language, code, theme, indent) {
  this.editor = editor;
  this.language = language;
  this.code = code;
  this.theme = theme;
  this.indent = indent;
}

Ground.prototype.setCursor = function() {
  var lastLine = this.editor.session.getLength();
  this.editor.gotoLine(lastLine);
  this.editor.focus();
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

Ground.prototype.setLanguage = function() {
  this.editor.getSession().setMode("ace/mode/" + this.language.code);
  $("#language-name").text(this.language.label);
};

Ground.prototype.cleanConsole = function() {
   $("#console").find("li").each(function() {
      this.remove();
   });
};

Ground.prototype.runCode = function() {
  data = JSON.stringify({ language: this.language, code: this.code});
  socket.send(data);
};

Ground.prototype.bindEditorEvents = function() {
  // Refresh language
  $(".language-link").on('click', function(event, date) {
    this.language = $(event.currentTarget).data('language');
    this.setLanguage();
  });
  // Refresh theme
  $(".theme-link").on('click', function(event, date) {
    this.theme = $(event.currentTarget).data('theme');
    this.setTheme();
  });
  // Refresh indentation
  $(".indent-link").on('click', function(event, date) {
    this.indent = $(event.currentTarget).data('indent');
    this.setIndent();
  });
  // Refresh code sample
  $(".language-link").on('ajax:complete', function(event, data) {
    if (data.status == 200) {
      response = JSON.parse(data.responseText);
      this.editor.setValue(response.custom);
      this.setCursor(editor);   
    } 
  });
  // Open socket to web server
  var socket = new WebSocket("ws://" + window.location.host + "/grounds/run");
  socket.onmessage = function(event) {
    if (event.data.length) {
      $("#console").append($('<li>').text(event.data));
    }
  };
  // Form submit
  $("#new_ground").on('submit', function(event) {
    event.preventDefault();
    this.cleanConsole();

    var code = this.editor.getValue();
    var language = $("#ground_language").val();
    this.runCode(socket, language, code);
  });
};
