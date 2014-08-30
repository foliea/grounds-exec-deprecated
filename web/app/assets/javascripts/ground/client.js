function Client(endpoint) {
  this.endpoint = endpoint;
  this.socket = null;
}

Client.prototype.connect = function() {
  if (!window.WebSocket) return false;
  this.socket = new WebSocket(this.endpoint);
  this.bindEvents();
  return true
};

Client.prototype.send = function(data) {
  if (this.socket === null) {
    var ok = this.connect();
    if (!ok) {
      $("#error").show();
      return;
    }
  }
  this.write(data);
};

Client.prototype.write = function(data) {
  var that = this;
  setTimeout(function(){
    if (that.socket === null) {
      $("#error").show();
      return;
    }
    switch (that.socket.readyState) {
      // CONNECTING
      case 0:
        that.write(data);
        break;
      // OPEN
      case 1:
        that.socket.send(data);
        break;
    }
  }, 1);
}

Client.prototype.bindEvents = function() {
  this.socket.onmessage = function(event) {
    $("#waiting").hide();
    if (!event.data.length) return;
    
    response = JSON.parse(event.data);
    if (response.stream === 'error') {
      $("#error").show();
      return;
    }
    if (response.stream === 'status') {
      response.chunk = "\n[Program exited with status: " + response.chunk + "]";
      $("body").animate({scrollTop:$(document).height()}, 1000);
    }
    $("#console").append($('<span class="'+ response.stream +'">').text(response.chunk));
  };
  var that = this;
  this.socket.onclose = function() {
    that.socket = null;
  };
  // Handle any errors that occur.
  this.socket.onerror = function(error) {
    console.log('WebSocket Error: ' + error);
  };
};
