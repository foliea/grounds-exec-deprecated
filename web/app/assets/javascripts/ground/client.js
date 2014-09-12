function Client(endpoint) {
    this.socket = null;
    this.connected = false;
 
    this.connect(endpoint);
    this.reload();
}

Client.prototype.reload = function() {
    this.console = new Console();
};

Client.prototype.connect = function(endpoint) {
    if (endpoint === null) return;
  
    this.socket = io.connect(endpoint);
    this.bindEvents();
};

Client.prototype.send = function(event, data) {
    if (this.connected === false) return;

    this.console.startWaiting();
     
    var request = JSON.stringify(data);
  
    this.socket.emit(event, request);
};

Client.prototype.bindEvents = function() {
    var that = this;
    this.socket.on('run', function(data) {
        var response = JSON.parse(data);
        that.console.write(response.stream, response.chunk);
    });
    this.socket.on('connect', function(data) {
        that.connected = true;
    });
    this.socket.on('connect_error', function() {
        that.connected = false;
        that.console.error();
    });
};
