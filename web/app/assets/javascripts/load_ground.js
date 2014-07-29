var loadGround = function() {
  // Return if there is no editor on the page
  var $groundEditor = $("#ground_editor");
  if (!$groundEditor[0]) {
    return;
  }
  // Load data
  var editor = ace.edit("ground_editor");
  var theme = $groundEditor.data("theme");
  var language = $groundEditor.data("language");
  var indent = $groundEditor.data("indent");
  var error = $groundEditor.data("error");
  
  // Create ground
  var ground = new Ground(editor, language, theme, indent);
};

$(document).ready(loadGround);
$(document).on("page:load", loadGround);
