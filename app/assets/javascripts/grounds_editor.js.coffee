loadGroundsEditor = ->
  editor = ace.edit('grounds_editor')
  editor.setTheme 'ace/theme/textmate'
  editor.getSession().setMode 'ace/mode/golang'

  $('#new_ground').submit ->
    editorContent = editor.getValue()
    $('#ground_code').val(editorContent)
  return

$(document).ready loadGroundsEditor
$(document).on 'page:load', loadGroundsEditor
