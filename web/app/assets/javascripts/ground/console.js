function Console() {
  $("#sharedURL").hide();
}

Console.prototype.clean = function() {
  $("#waiting").hide();
  $("#error").hide();
  $("#console").find("span").each(function() {
    this.remove();
  });
};

Console.prototype.startWaiting = function() {
  this.clean();
  $("#waiting").show();
};

Console.prototype.stopWaiting = function() {
  $("#waiting").hide();
};

Console.prototype.write = function(stream, chunk) {
  if (stream === 'status') {
    if (chunk !== 0)
      $("#console").find("span").each(function() {
        $(this).removeClass('stdout')
               .addClass('stderr');
      });

    chunk = '[Program exited with status: ' + chunk + ']';
  }
  $("#console").append($('<span class="' + stream + '">').text(chunk));
};

Console.prototype.error = function() {
  this.stopWaiting();
  $("#error").show();
};
