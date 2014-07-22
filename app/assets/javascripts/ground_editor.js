var loadGroundEditor = function() {
  var $groundEditor = $("#ground_editor");
  if (!$groundEditor[0]) {
    return;
  }

  var theme = $groundEditor.data("theme");
  var language = $groundEditor.data("language");
  var error = $groundEditor.data("error");

  var editor = ace.edit("ground_editor");
  editor.setTheme("ace/theme/" + theme);
  editor.getSession().setMode("ace/mode/" + language);
 
  // Set cursor on the last line
  editor.gotoLine(editor.session.getLength());
  editor.focus();

  $("#new_ground").submit(function() {
    $("#ground_code").val(editor.getValue());
  });
 
  $("#new_ground").on("ajax:complete", function(event, data) {
    if (data.status == 200) {
      output = JSON.parse(data.responseText);
      $("#stdout").text(output.stdout);
      $("#stderr").text(output.stderr);
    } else {
      $("#stdout").text('');
      $("#stderr").text(error);
    }
  });
  
  editor.commands.addCommand({
    name: 'Undo',
    bindKey: {win: 'Ctrl-Z',  mac: 'Command-Z'},
    exec: function(editor) {
      editor.undo();
    },
    readOnly: true
  });
};

$(document).ready(loadGroundEditor);
$(document).on("page:load", loadGroundEditor);
