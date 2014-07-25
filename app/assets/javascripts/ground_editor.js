var loadGroundEditor = function() {
  var setTheme = function(editor, theme) {
    editor.setTheme("ace/theme/" + theme.code);
    $("#theme-name").text(theme.label);
  };
  
  var setLanguage = function(editor, language) {
    editor.getSession().setMode("ace/mode/" + language);
  }
  
  var setCursor = function(editor) {
    // Set cursor on the last line
    editor.gotoLine(editor.session.getLength());
    editor.focus();
  }

  var bindCommands = function(editor) {
    editor.commands.addCommand({
      name: 'Undo',
      bindKey: {win: 'Ctrl-Z',  mac: 'Command-Z'},
      exec: function(editor) {
        editor.undo();
      },
      readOnly: true
    });
  }
  
  var bindFormEvents = function(editor) {
    // Form submission
    $("#new_ground").submit(function() {
      $("#ground_code").val(editor.getValue());
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
  
  var bindThemeEvents = function(editor) {
    $(".theme-link").on("click", function(event, date) {
      var theme = $(event.currentTarget).data('theme');
      setTheme(editor, theme);
    });
  };

  var $groundEditor = $("#ground_editor");
  if (!$groundEditor[0]) {
    return;
  }
  var theme = $groundEditor.data("theme");
  var language = $groundEditor.data("language");
  var error = $groundEditor.data("error");
  var editor = ace.edit("ground_editor");

  setTheme(editor, theme);
  setLanguage(editor, language);
  setCursor(editor);
  bindCommands(editor);
  bindFormEvents(editor);
  bindThemeEvents(editor);
};

$(document).ready(loadGroundEditor);
$(document).on("page:load", loadGroundEditor);
