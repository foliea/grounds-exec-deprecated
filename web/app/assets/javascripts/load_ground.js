var ground = null;

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
  var keyboard = $groundEditor.data("keyboard");
  var dockerUrl = $groundEditor.data("docker-url");

  var runner = new Runner(dockerUrl);
  ground = new Ground(editor, language, theme, indent, keyboard, runner);
};

$(document).ready(loadGround);
$(document).on("page:load", loadGround);
