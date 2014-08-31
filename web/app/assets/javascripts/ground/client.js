function Client(endpoint) {
  this.endpoint = endpoint;
  this.socket = null;
}

Client.prototype.connect = function() {
  this.socket = io.connect('http://127.0.0.1:5000');
  this.bindEvents();
};

Client.prototype.send = function(data) {
  this.socket.emit('run message', data);
};

Client.prototype.bindEvents = function() {
  this.socket.on('run message', function(data) {
    alert(JSON.parse(data));
  });
};
