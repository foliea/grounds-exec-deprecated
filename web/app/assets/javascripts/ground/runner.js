function Runner(dockerUrl) {
  this.dockerUrl = dockerUrl;
  this.console = new Console();
  this.previousContainer = null;
}

Runner.prototype.run = function(language, code) {
    if (this.previousContainer)
      this.previousContainer.interrupt();
    this.previousContainer = new Container(this.dockerUrl, this.console);
    this.previousContainer.run(language, code)
};
