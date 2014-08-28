module GroundControls
  GROUND = 'ground'
  
  def refresh
    visit(ground_show_path)
  end

  def show_dropdown(option)
    find("a[data-dropdown='#{option.pluralize(2)}']").click
  end

  def select_option(option, code)
    find("a[data-#{option}='#{code}']").click
  end
  
  def share
    find('#share').click
  end
  
  def type_inside_editor
    evaluate_script("#{GROUND}.editor.setValue('typing...');")
  end

  def shared_url
    find('input[name="sharedURL"]').value
  end

  def selected_option_label(option, code)
    find("##{option}-name").text
  end
  
  def session(option)
    page.get_rack_session_key(option)
  end
  
  def default_option_code(option)
    GroundEditor.default_option_code(option)
  end

  def option_label(option, code)
    GroundEditor.option(option, code)[:label]
  end
  
  def editor_content
    find('#ground_editor').text
  end
  
  def editor_mode
    mode = evaluate_script("#{GROUND}.editor.getSession().getMode().$id;")
    mode.gsub('ace/mode/', '')
  end
  
  def editor_code
    evaluate_script("#{GROUND}.editor.getValue();")
  end
  
  def editor_cursor_on_last_line?
    pos = evaluate_script("#{GROUND}.editor.getCursorPosition();")
    line = evaluate_script("#{GROUND}.editor.session.getLength();") - 1;
    pos['row'].to_i == line
  end
  
  def editor_theme
    theme = evaluate_script("#{GROUND}.editor.getTheme();")
    theme.gsub('ace/theme/', '')
  end

  def editor_tab_size
    evaluate_script("#{GROUND}.editor.getSession().getTabSize();")
  end
  
  def editor_use_soft_tabs?
    evaluate_script("#{GROUND}.editor.getSession().getUseSoftTabs();")
  end

  def mode(language)
    evaluate_script("GetMode('#{language}');")
  end
  
  def sample(language)
    evaluate_script("GetSample('#{language}');")
  end
end