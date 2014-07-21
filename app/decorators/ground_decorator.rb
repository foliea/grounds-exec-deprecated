class GroundDecorator < BaseDecorator
  def editor
    {
      theme: h.session[:editor_theme],
      language:  self.language
	  }
  end
end