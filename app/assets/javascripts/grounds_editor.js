var loadGroundsEditor = function() {
  var groundsEditor = $("#grounds_editor");
  if (!groundsEditor[0]) {
    return;
  }

  var theme = groundsEditor.data("theme");
  var language = groundsEditor.data("language");
  var editor = ace.edit("grounds_editor");

  editor.setTheme("ace/theme/" + theme);
  editor.getSession().setMode("ace/mode/" + language);
 
  $("#new_ground").submit(function() {
    var editorContent = editor.getValue();
    $("#ground_code").val(editorContent);
  });
 
  $("#new_ground").on("ajax:complete", function(event, data) {
    if (data.status == 200) {
      output = JSON.parse(data.responseText);
      $("#stdout").text(output.stdout);
      $("#stderr").text(output.stderr);
    }
  });
};

$(document).ready(loadGroundsEditor);
$(document).on("page:load", loadGroundsEditor);
