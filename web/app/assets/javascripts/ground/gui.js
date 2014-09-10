function GUI(ground, client) {
    this._ground = ground;
    this._client = client;

    this.sharedURL = $("#sharedURL");

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

    this.initialize();
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

GUI.prototype.dropdownSelect = function(dropdown, value) {
    $('a[data-dropdown="' + dropdown + 's"]').click();
    $("#" + dropdown + "-name").text(value);
};

GUI.prototype.initialize = function() {
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
        
        // Move to client
        var data = JSON.stringify({ language: language, code: code });
    
        that._client.send(data);
    });
    
    this.button.back.on('click', function(event) {
        $("body").animate({scrollTop: 0}, 'fast');
        that._ground._editor.focus(); 
    });
    
    this.link.language.on('click', function(event, date) {
        var link = $(this);
        var language = link.data('language');
        var label = link.text();

        that._ground.setLanguage(language);
        that.dropdownSelect('language', label);
    });
    
    this.link.theme.on('click', function(event, date) {
        var link = $(this);
        var theme = link.data('theme');
        var label = link.text();

        that._ground.setTheme(theme);
        that.dropdownSelect('theme', label);
    });
    
    this.link.indent.on('click', function(event, date) {
        var link = $(this);
        var indent = link.data('indent');
        var label = link.text();

        that._ground.setIndent(indent);
        that.dropdownSelect('indent', label);
    });
    
    this.link.keyboard.on('click', function(event, date) {
        var link = $(this);
        var keyboard = link.data('keyboard');
        var label = link.text();

        that._ground.setKeyboard(keyboard);
        that.dropdownSelect('keyboard', label);
    });
    
    this.form.obj.on("ajax:success", function(data, response, xhr) {
        if (response.status !== 'ok') return;
        that.sharedURL.val(response.shared_url).show().focus().select();
    });

    this._ground._editor.on('input', function() {
        that.sharedURL.hide();
    });
}