var loadGroundEditor = function() {
  var setTheme = function(editor, theme) {
    editor.setTheme("ace/theme/" + theme.code);
    $("#theme-name").text(theme.label);
  };
  
  var setIndent = function(editor, indent) {
    if (indent.code == "tab") {
      editor.getSession().setUseSoftTabs(false);
      editor.getSession().setTabSize(8);
    } else {
      editor.getSession().setUseSoftTabs(true);
      editor.getSession().setTabSize(indent.code);
    }
    $("#indent-name").text(indent.label);
  };
  
  var setLanguage = function(editor, language) {
    editor.getSession().setMode("ace/mode/" + language.code);
    $("#language-name").text(language.label);
  }
  
  var setCursor = function(editor) {
    // Set cursor on the last line
    editor.gotoLine(editor.session.getLength());
    editor.focus();
  }
  
  var bindFormEvents = function(editor) {
    // Form submission
    $("#new_ground").submit(function() {
      $("#ground_code").val(editor.getValue());
      // Clean console
      $("#stdout").text('');
      $("#stderr").text('');
    });
    // Get response after form submission
    $("#new_ground").on("ajax:complete", function(event, data) {
      if (data.status == 200) {
        response = JSON.parse(data.responseText);
        $("#stdout").text(response.stdout);
        $("#stderr").text(response.stderr);
      } else {
        $("#stdout").text('');
        $("#stderr").text(error);
      }
    });
  };
  
  var bindEditorEvents = function(editor) {
    $(".language-link").on("click", function(event, date) {
      var language = $(event.currentTarget).data('language');
      setLanguage(editor, language);
    });
    $(".theme-link").on("click", function(event, date) {
      var theme = $(event.currentTarget).data('theme');
      setTheme(editor, theme);
    });
    $(".indent-link").on("click", function(event, date) {
      var indent = $(event.currentTarget).data('indent');
      setIndent(editor, indent);
    });
  };

  var $groundEditor = $("#ground_editor");
  if (!$groundEditor[0]) {
    return;
  }
  var theme = $groundEditor.data("theme");
  var indent = $groundEditor.data("indent");
  var language = $groundEditor.data("language");
  var error = $groundEditor.data("error");
  var editor = ace.edit("ground_editor");

  setLanguage(editor, language);
  setTheme(editor, theme);
  setIndent(editor, indent);
  setCursor(editor);
  bindFormEvents(editor);
  bindEditorEvents(editor);
  
  editor.getSession().setUseWrapMode(true);
};

$(document).ready(loadGroundEditor);
$(document).on("page:load", loadGroundEditor);
