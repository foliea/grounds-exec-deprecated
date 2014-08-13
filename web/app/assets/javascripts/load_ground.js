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
  var run_endpoint = "ws://" + $groundEditor.data("run-endpoint");
 
  // Create ground
  var client = new Client(run_endpoint);

  // FIXME: WEBSOCKET non existent
  var ground = new Ground(editor, language, theme, indent, client);
};

$(document).ready(loadGround);
$(document).on("page:load", loadGround);
