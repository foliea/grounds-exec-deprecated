class GroundDecorator < BaseDecorator
  def editor
    {
      theme: h.session[:theme] ||= default_theme,
      indent: h.session[:indent] ||= default_indent,
      language:  self.language,
      error:  'An error occured, please try again later.'
	  }
  end
  
  def default_theme
    code, label = GroundEditor.themes.first
    { code: code, label: label}
  end
  
  def default_indent
    code, label = GroundEditor.indents.first
    { code: code, label: label}
  end
end


