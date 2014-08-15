// FIXME: Alert if browser doesn't support websockets
// FIXME: Do something if connection to websockets is impossible
// FIXME: readyState3
function Client(endpoint) {
  this.endpoint = endpoint;
  this.socket = null;
}

Client.prototype.connect = function() {
  this.socket = new WebSocket(this.endpoint);
  this.bindEvents();
}

Client.prototype.bindEvents = function() {
  this.socket.onmessage = function(event) {
    $("#waiting").hide();
    if (event.data.length) {
      response = JSON.parse(event.data);
      if (response.stream === 'status') {
        response.chunk = "\n[Program exited with status: " + response.chunk + "]";
        $("body").animate({scrollTop:$(document).height()}, 1000);
      }
      $("#console").append($('<span class="'+ response.stream +'">').text(response.chunk));
    }
  };
  var that = this;
  this.socket.onclose = function() {
    that.socket = null;
  };
}
// FIXME: stop connection attempt if 10 fails
Client.prototype.send = function(data) {
  $("#waiting").show();
  if (this.socket === null) {
    this.connect(); 
  }
  var that = this;
  setTimeout(function(){
    if (that.socket !== null && that.socket.readyState === 1) {
      that.socket.send(data);
      return;
    } else {
      that.send(data); 
    }
  }, 1);
};
