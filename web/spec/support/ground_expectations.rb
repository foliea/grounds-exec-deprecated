module GroundExpectations
  def expect_data(option, code)
    data = find("#ground_editor[data-#{option}='#{code}']")
    expect(data).not_to be_nil
  end

  def expect_selected_label(option, code)
    label = selected_option_label(option, code)
    expect(label).to eq(option_label(option, code))
  end

  def expect_shared_url_visibility(value)
    expect(find('#sharedURL', visible: value)).not_to be_nil
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
    expect(editor_mode).to eq(mode(language))
    expect(editor_code).to eq(sample(language))
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
end
