function Client(endpoint) {
    this._socket = null;
    this._connected = false;
 
    this.connect(endpoint);
    this.reload();
}

Client.prototype.reload = function() {
    this._console = new Console();
};

Client.prototype.connect = function(endpoint) {
    if (endpoint === null) return;
  
    this._socket = io.connect(endpoint);
    this.bindEvents();
};

Client.prototype.send = function(event, data) {
    if (this._connected === false) return;

    this._console.startWaiting();
     
    var request = JSON.stringify(data);
  
    this._socket.emit(event, request);
};

Client.prototype.bindEvents = function() {
    var that = this;
    this._socket.on('run', function(data) {
        var response = JSON.parse(data);
        that._console.write(response.stream, response.chunk);
    });
    this._socket.on('connect', function(data) {
        that._connected = true;
        that._console.clean();
    });
    this._socket.on('connect_error', function() {
        that._connected = false;
        that._console.error();
    });
};
