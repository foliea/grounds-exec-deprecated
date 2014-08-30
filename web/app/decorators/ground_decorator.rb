class GroundDecorator < BaseDecorator
  def editor
    {
      theme: h.session[:theme] ||= GroundEditor.default_option_code(:theme),
      indent: h.session[:indent] ||= GroundEditor.default_option_code(:indent),
      keyboard: h.session[:keyboard] ||= GroundEditor.default_option_code(:keyboard),
      language: self.language
    }
  end

  def themes
    GroundEditor.options(:theme)
  end

  def indents
    GroundEditor.options(:indent)
  end
  
  def languages
    GroundEditor.options(:language)
  end
  
  def keyboards
    GroundEditor.options(:keyboard)
  end
end


