class GroundDecorator < BaseDecorator
  def editor
    {
      theme: h.session[:editor_theme] || 'textmate',
      language:  self.language,
      error:  'An error occured, please try again later.'
	  }
  end
end