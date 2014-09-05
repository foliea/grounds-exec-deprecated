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
  this.console.startWaiting();
  
  if (this.socket === null) {
    this.console.error();
    return;
  }
  this.socket.emit('run', data);
};

Client.prototype.bindEvents = function() {
  var that = this;
  this.socket.on('run', function(data) {
    var response = JSON.parse(data);
    that.console.write(response.stream, response.chunk);
  });
};
