module GroundExpectations
  def expect_data(option, code)
    data = find("#ground_editor[data-#{option}='#{code}']")
    expect(data).not_to be_nil
  end

  def expect_selected_label(option, code)
    label = selected_option_label(option, code)
    expect(label).to eq(option_label(option, code))
  end
  
  def expect_option_in_session(option, code)
    session_option = page.get_rack_session_key(option)
    expect(session_option).to eq(code)
  end

  def selected_option_label(option, code)
    find("##{option}-name").text
  end

  def option_label(option, code)
    GroundEditor.option(option, code)[:label]
  end
  
  def expect_editor_option(option, code)
    case option
    when 'language'
      expect_editor_language(code)
    when 'theme'
      expect_editor_theme(code)
    when 'indent'
      expect_editor_indent(code)
    end
  end
  
  def expect_editor_language(language)
    expect(editor_mode).to eq(to_mode(language))
    expect(editor_code).to eq(to_sample(language))
    expect(editor_cursor_on_last_line?).to be true
  end
  
  def expect_editor_theme(theme)
    expect(editor_theme).to eq(theme)
  end
  
  def expect_editor_indent(indent)
    use_soft_tabs = indent == 'tab' ? false : true
    expect(editor_use_soft_tabs?).to eq(use_soft_tabs)

    tab_size = indent == 'tab' ? 8 : indent.to_i
    expect(editor_tab_size).to eq(tab_size)
  end

  def editor_mode
    mode = evaluate_script('ground.editor.getSession().getMode().$id;')
    mode.gsub('ace/mode/', '')
  end
  
  def editor_code
    evaluate_script('ground.editor.getValue();')
  end
  
  def editor_cursor_on_last_line?
    pos = evaluate_script('ground.editor.getCursorPosition();')
    line = evaluate_script('ground.editor.session.getLength();') - 1;
    pos['row'].to_i == line
  end
  
  def editor_theme
    theme = evaluate_script('ground.editor.getTheme();')
    theme.gsub('ace/theme/', '')
  end

  def editor_tab_size
    evaluate_script('ground.editor.getSession().getTabSize();')
  end
  
  def editor_use_soft_tabs?
    evaluate_script('ground.editor.getSession().getUseSoftTabs();')
  end

  def to_mode(language)
    evaluate_script("GetMode('#{language}');")
  end
  
  def to_sample(language)
    evaluate_script("GetSample('#{language}');")
  end
end
