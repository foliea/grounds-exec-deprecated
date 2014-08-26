module GroundExpectations
  def expect_data(option, code)
    data = find("#ground_editor[data-#{option}=#{code}]")
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
end
