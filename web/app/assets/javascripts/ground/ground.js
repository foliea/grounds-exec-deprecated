function Ground(editor, language, theme, indent, keyboard, client) {
  this.editor = editor;
  this.language = language;
  this.theme = theme;
  this.indent = indent;
  this.keyboard = keyboard;
  this.client = client;

  this.initEditor();
  this.setLanguage();
  this.setTheme();
  this.setIndent();
  this.setKeyboard();

  this.bindEvents();

  this.client.connect();
}

Ground.prototype.initEditor = function() {
  this.keybindings = {
    ace: null, // use "default" keymapping
    vim: "ace/keyboard/vim",
    emacs: "ace/keyboard/emacs"
  };
};

Ground.prototype.setCursor = function() {
  var lastLine = this.editor.session.getLength();
  this.editor.gotoLine(lastLine);
  this.editor.focus();
};

Ground.prototype.setLanguage = function() {
  this.editor.getSession().setMode("ace/mode/" + GetMode(this.language));
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

Ground.prototype.setKeyboard = function() {
  this.editor.setKeyboardHandler(this.keybindings[this.keyboard]);
};

Ground.prototype.bindEvents = function() {
  var that = this;
  // Refresh language
  $(".language-link").on('click', function(event, date) {
    that.language = $(event.currentTarget).data('language');
    that.editor.setValue("");
    that.setLanguage();
    $('a[data-dropdown="languages"]').click();
    $("#language-name").text($(this).text());
  });
  // Refresh theme
  $(".theme-link").on('click', function(event, date) {
    that.theme = $(event.currentTarget).data('theme');
    that.setTheme();
    $('a[data-dropdown="themes"]').click();
    $("#theme-name").text($(this).text());
  });
  // Refresh indentation
  $(".indent-link").on('click', function(event, date) {
    that.indent = $(event.currentTarget).data('indent');
    that.setIndent();
    $('a[data-dropdown="indents"]').click();
    $("#indent-name").text($(this).text());
  });
  // Refresh keyboard
  $(".keyboard-link").on('click', function(event, date) {
    that.keyboard = $(event.currentTarget).data('keyboard');
    that.setKeyboard();
    $('a[data-dropdown="keyboards"]').click();
    $("#keyboard-name").text($(this).text());
  }); 
  // Form submit
  $("#run").on('click', function(event) {
    var code = that.editor.getValue();
    var language = that.language;
    var data = JSON.stringify({ language: language, code: code });
    that.client.send(data);
  });
  // Share current snippet
  $("#share").on('click', function(event) {
    var code = that.editor.getValue();
    $("#ground_code").val(code);
    $("#ground_language").val(that.language);
    $("#share_ground").submit();
  });
  // Scroll back to editor
  $("#back").on('click', function(event) {
    $("body").animate({scrollTop: 0}, 'fast');
    that.editor.focus(); 
  });
  // Get result from share action and display shared link
  $("#share_ground").on("ajax:success", function(data, response, xhr) {
    if (response.status !== "ok") {
      that.client.console.clean();
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
