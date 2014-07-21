loadGroundsEditor = ->
  groundsEditor = $('#grounds_editor')
  return  unless groundsEditor[0]
  
  theme = groundsEditor.data('theme')
  language = groundsEditor.data('language')
  
  editor = ace.edit('grounds_editor')
  editor.setTheme 'ace/theme/' + theme
  editor.getSession().setMode 'ace/mode/' + language

  $('#new_ground').submit ->
    editorContent = editor.getValue()
    $('#ground_code').val(editorContent)
  return

$(document).ready loadGroundsEditor
$(document).on 'page:load', loadGroundsEditor
