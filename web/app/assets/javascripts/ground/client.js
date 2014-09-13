function Client(endpoint) {
    this._console = new Console();
    this._socket = null;

    this.connect(endpoint);
}

Client.prototype.connect = function(endpoint) {
    if (endpoint === null) return;

    this._socket = io.connect(endpoint, {'forceNew':true });
    this.bindEvents();
};

Client.prototype.disconnect = function() {
    if (this.connected() === false) return;

    this._socket.io.disconnect();
};

Client.prototype.connected = function() {
    return this._socket !== null && this._socket.connected;
};

Client.prototype.send = function(event, data) {
    if (this.connected() === false) return;

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
        that._console.clean();
    });
    this._socket.on('connect_error', function() {
        that._console.error();
    });
};
