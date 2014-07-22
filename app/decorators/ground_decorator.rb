class GroundDecorator < BaseDecorator
  def editor
    {
      theme: h.session[:editor_theme] || 'tomorrow_night',
      language:  self.language,
      error:  'An error occured, please try again later.'
	  }
  end
end