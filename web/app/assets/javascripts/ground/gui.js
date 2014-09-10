function GUI(ground, client) {
    this._ground = ground;
    this._client = client;

    this.sharedURL = $("#sharedURL").hide();

    this.button = {
        share: $("#share"),
        run: $("#run"),
        back: $("#back"),
    };
    
    this.form = {
        obj: $("#share_ground"),
        code: $("#ground_code"),
        language: $("#ground_language"),
    };
    
    this.link = {
        language: $(".language-link"),
        theme: $(".theme-link"),
        indent: $(".indent-link"),
        keyboard: $(".keyboard-link"),
    }

    this.bindEvents();
}

GUI.prototype.submitShareFormWith = function(language, code) {
    this.form.language.val(language);
    this.form.code.val(code);
    this.form.obj.submit();
};

GUI.prototype.disableRunButtonFor = function(milliseconds) {
    this.button.run.attr('disabled', 'disabled');

    var that = this;
    setTimeout(function() {
      that.button.run.removeAttr('disabled');
    }, milliseconds);
};

GUI.prototype.dropdownUpdate = function(option, label) {
    $('a[data-dropdown="' + option + 's"]').click();
    $("#" + option + "-name").text(label);
};

GUI.prototype.scrollToTop = function() {
    $("body").animate({scrollTop: 0}, 'fast');
};

GUI.prototype.switchToSelectedOption = function(option, link) {
    var code = link.data(option);
    var label = link.text();
  
    this._ground.set(option, code);
    this.dropdownUpdate(option, label);
};

GUI.prototype.bindEvents = function() {
    var that = this;
    this.button.share.on('click', function(event) {
        var language = that._ground.getLanguage();
        var code = that._ground.getCode();
      
        that.submitShareFormWith(language, code);
    });
    
    this.button.run.on('click', function(event) {
        that.disableRunButtonFor(500);
    
        var language = that._ground.getLanguage();
        var code = that._ground.getCode();
    
        that._client.send('run', { language: language, code: code });
    });
    
    this.button.back.on('click', function(event) {
        that.scrollToTop();
        that._ground._editor.focus(); 
    });
    
    this.link.language.on('click', function(event, date) {
        that.switchToSelectedOption('language', $(this));
    });
    
    this.link.theme.on('click', function(event, date) {
        that.switchToSelectedOption('theme', $(this));
    });
    
    this.link.indent.on('click', function(event, date) {
        that.switchToSelectedOption('indent', $(this));
    });
    
    this.link.keyboard.on('click', function(event, date) {
        that.switchToSelectedOption('keyboard', $(this));
    });
    
    this.form.obj.on('ajax:success', function(data, response, xhr) {
        if (response.status !== 'ok') return;
        that.sharedURL.val(response.shared_url).show().focus().select();
    });

    this._ground._editor.on('input', function() {
        that.sharedURL.hide();
    });
};