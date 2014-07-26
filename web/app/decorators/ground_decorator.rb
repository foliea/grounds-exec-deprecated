class GroundDecorator < BaseDecorator
  def editor
    {
      theme: h.session[:theme] ||= GroundEditor.default_option(:theme),
      indent: h.session[:indent] ||= GroundEditor.default_option(:indent),
      language:  self.language,
      error:  'An error occured, please try again later.'
	  }
  end
end


