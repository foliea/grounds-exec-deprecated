class GroundDecorator < BaseDecorator
  def editor
    {
      theme: h.session[:editor_theme] ||= Theme.all.first,
      language:  self.language,
      error:  'An error occured, please try again later.'
	  }
  end
  
  def themes
    Theme.all
  end
end

module Theme
  extend self

  def all
    [
      { label: 'Textmate', code: 'textmate' } ,
      { label: 'Monokai', code: 'monokai' },
      { label: 'Tomorrow Night', code: 'tomorrow_night' } 
    ]
  end
  
  def get(code)
    all.each { |theme| return theme if theme[:code] == code  }
  end
end