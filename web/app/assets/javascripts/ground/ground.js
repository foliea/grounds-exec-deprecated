function Ground(editor, language, theme, indent, client) {
  this.editor = editor;
  this.language = language;
  this.theme = theme;
  this.indent = indent;
  this.client = client;

  this.initEditor();
  this.setLanguage();
  this.setTheme();
  this.setIndent();

  this.bindEvents();

  this.editor.commands.addCommand({
    name: 'Run',
    bindKey: {win: 'Ctrl-K',  mac: 'Command-K'},
    exec: function(editor) {
        $("#run").click();
    },
    readOnly: false
  });
  this.editor.commands.addCommand({
    name: 'Back to editor',
    bindKey: {win: 'Ctrl-J',  mac: 'Command-J'},
    exec: function(editor) {
        $("#back").click();
    },
    readOnly: false
  });
}

Ground.prototype.initEditor = function() {
  $("#sharedURL").hide();
  this.editor.getSession().setUseWrapMode(true);
};

Ground.prototype.setCursor = function() {
  var lastLine = this.editor.session.getLength();
  this.editor.gotoLine(lastLine);
  this.editor.focus();
};

Ground.prototype.setLanguage = function() {
  this.editor.getSession().setMode("ace/mode/" + GetTheme(this.language));
  if (this.editor.getValue() === "")
    this.editor.setValue(GetSample(this.language));
  this.setCursor();
};

Ground.prototype.setTheme = function() {
  this.editor.setTheme("ace/theme/" + this.theme);
};

Ground.prototype.setIndent = function() {
  if (this.indent == "tab") {
    this.editor.getSession().setUseSoftTabs(false);
    this.editor.getSession().setTabSize(8);
  } else {
    this.editor.getSession().setUseSoftTabs(true);
    this.editor.getSession().setTabSize(this.indent);
  }
};

Ground.prototype.cleanConsole = function() {
  $("#waiting").hide();
   $("#error").hide();
   $("#console").find("span").each(function() {
      this.remove();
   });
};

Ground.prototype.bindEvents = function() {
  var that = this;
  // Refresh language
  $(".language-link").on('click', function(event, date) {
    that.language = $(event.currentTarget).data('language');
    that.editor.setValue("");
    that.setLanguage();
    $("#language-name").text($(this).text());
  });
  // Refresh theme
  $(".theme-link").on('click', function(event, date) {
    that.theme = $(event.currentTarget).data('theme');
    that.setTheme();
    $("#theme-name").text($(this).text());
  });
  // Refresh indentation
  $(".indent-link").on('click', function(event, date) {
    that.indent = $(event.currentTarget).data('indent');
    that.setIndent();
    $("#indent-name").text($(this).text());
  }); 
  // Form submit
  $("#run").on('click', function(event) {
    that.cleanConsole();
    $("#waiting").show();

    var code = that.editor.getValue();
    var language = that.language;
    data = JSON.stringify({ language: language, code: code });
    that.client.send(data);
  });
  // Share current snippet
  $("#share").on('click', function(event) {
    var code = that.editor.getValue();
    $("#ground_code").val(code);
    $("#ground_language").val(that.language);
    $("#new_ground").submit();
  });
  // Scroll back to editor
  $("#back").on('click', function(event) {
    $("body").animate({scrollTop: 0}, 'fast');
    that.editor.focus(); 
  });
  // Get result from share action and display shared link
  $("#new_ground").on("ajax:success", function(data, response, xhr) {
    if (response.status !== "ok") {
      that.cleanConsole();
      $("#error").show();
      return;
    }
    var sharedURL = response.shared_url;
    $("#sharedURL").val(sharedURL)
                   .show()
                   .focus()
                   .select();
  });
  // Hide shared url if code is modified
  that.editor.on('input', function() {
    $("#sharedURL").hide();
  });
};
