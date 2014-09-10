function Console() {
  this.output = $("#console");
  this.error = $("#error");
  this.waiting = $("#waiting");
}

Console.prototype.clean = function() {
  this.waiting.hide();
  this.error.hide();
  this.output.find("span").each(function() {
    this.remove();
  });
};

Console.prototype.startWaiting = function() {
  this.clean();
  this.waiting.show();
};

Console.prototype.stopWaiting = function() {
  this.waiting.hide();
};

Console.prototype.write = function(stream, chunk) {
  switch(stream) {
    case "status":
      this.stopWaiting();
      chunk = '[Program exited with status: ' + chunk + ']';
      break;
    case "error":
      this.stopWaiting();
      stream = 'stderr';
      break;
  }
  this.output.append($('<span class="' + stream + '">').text(chunk));
};

Console.prototype.error = function() {
  this.stopWaiting();
  this.error.show();
};
