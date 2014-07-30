class GroundDecorator < BaseDecorator
  def editor
    {
      theme: h.session[:theme] ||= GroundEditor.default_option(:theme),
      indent: h.session[:indent] ||= GroundEditor.default_option(:indent),
      language:  self.language,
      error: I18n.t('editor.error') 
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
end


