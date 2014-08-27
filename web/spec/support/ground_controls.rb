module GroundControls
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

  def shared_url
    find('input[name="sharedURL"]').value
  end
  
  def shared_link_visible?
    evaluate_script('$("#sharedURL").is(":visible");')
  end
  
  def type_inside_editor
    evaluate_script('ground.editor.setValue("typing...");')
  end
end