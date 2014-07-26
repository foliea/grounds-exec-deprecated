module ExecCode
  module Language
    extend self

    def all
      [
        'golang',
        'ruby'
      ]
    end

    def language_supported?(name)
      all.include?(name)
    end
  end 
end
