class Ground
  attr_reader :language
  attr_reader :code

  def initialize(language = nil)
    @language = language || GroundEditor.default_option(:language)
    @code = ExecCode::Sample.from(@language[:code]) || ''
  end

  def valid?
    GroundEditor.has_option?(:language, @language[:code])
  end
 end
