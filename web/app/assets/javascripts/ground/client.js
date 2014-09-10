function Client(endpoint) {
    this.socket = null;
    this.console = new Console();
  
    this.connect(endpoint);
}

Client.prototype.connect = function(endpoint) {
    if (endpoint === null) return;
  
    this.socket = io.connect(endpoint);
    this.bindEvents();
};

Client.prototype.send = function(event, data) {
    this.console.startWaiting();
    
    if (this.socket === null) {
        this.console.error();
        return;
    }
    var request = JSON.stringify(data);
  
    this.socket.emit(event, request);
};

Client.prototype.bindEvents = function() {
    var that = this;
    this.socket.on('run', function(data) {
        var response = JSON.parse(data);
        that.console.write(response.stream, response.chunk);
    });
};
