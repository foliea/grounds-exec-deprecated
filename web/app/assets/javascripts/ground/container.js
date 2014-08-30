function Container(dockerUrl, console) {
  this.dockerUrl = dockerUrl;
  this.console = console;

  this.id = null;
  this.socket = null;
  this.interrupted = false;
}

Container.prototype.run = function(language, code) {
  this.console.startWaiting(); 
  var that = this;
  var creation = this.create(language, code);
  creation.done(function(containerId) {
    that.id = containerId;
    that.attach();
    that.start(); 
  });
  creation.fail(function(err) {
    that.console.error();
  });
};

Container.prototype.interrupt = function() {
  if (this.interrupted) return;
  this.interrupted = true;
  this.stop();
};

Container.prototype.create = function(language, code) {
  var deferred = $.Deferred();
  var request = $.ajax({
    url: "/containers/create",
    type: "POST",
    data: { language: language, code: code }
  });
  request.always(function(response) {
    if (response.status === 'ok')
      deferred.resolve(response.container_id);
    else
      deferred.reject("HTTP error: " + response.status);
  });
  return deferred.promise();
};

Container.prototype.attach = function() {
  var that = this;
  var endpoint = formatEndpoint(this.dockerUrl, this.id);
  
  this.socket = new WebSocket(endpoint);
  this.socket.onopen = function() {
    that.socket.onmessage = function(message) {
      if (that.interrupted) return;
      that.console.stopWaiting();
      that.console.write("stdout", message.data);  
    };
    that.socket.onclose = function(message) {
      if (that.interrupted) return;
      that.interrupted = true;
      that.status();
    };  
  }; 
};

Container.prototype.start = function() {
  $.ajax({ 
    url: "/containers/start",
    type: 'POST', 
    data: { id: this.id },
    fail: this.console.error
  }); 
};

Container.prototype.stop = function() {
  $.ajax({ 
    url: "/containers/stop",
    type: 'POST', 
    data: { id: this.id },
    fail: this.console.error
  });
};

Container.prototype.status = function() {
  var that = this;
  request = $.ajax({
    url: "/containers/status", 
    data: { id: this.id } 
  });
  request.always(function(response) {
    if (response.status !== 'ok') return;
    that.console.write('status', response.code);
  });
};
