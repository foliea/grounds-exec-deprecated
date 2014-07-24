class GroundDecorator < BaseDecorator
  def editor
    {
      theme: h.session[:editor_theme] ||= themes.first,
      language:  self.language,
      error:  'An error occured, please try again later.'
	  }
  end
  
  def themes
    [
      { label: 'Textmate', code: 'textmate' } ,
      { label: 'Monokai', code: 'monokai' },
      { label: 'Tomorrow Night', code: 'tomorrow_night' } 
    ]
  end
end