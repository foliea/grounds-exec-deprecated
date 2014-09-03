function Client(endpoint) {
  this.endpoint = endpoint;
  this.socket = null;
  this.console = new Console();
}

Client.prototype.connect = function() {
  if (this.endpoint === null) return;

  this.socket = io.connect(this.endpoint);
  this.bindEvents();
};

Client.prototype.send = function(data) {
  if (this.socket === null) return;
  this.console.startWaiting();
  this.socket.emit('run', data);
};

Client.prototype.bindEvents = function() {
  var that = this;
  this.socket.on('run', function(data) {
    var response = JSON.parse(data);
    that.console.write(response.stream, response.chunk);
  });
};
