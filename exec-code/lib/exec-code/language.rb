module ExecCode
  module Language
    extend self

    def all
      [
        'golang',
        'ruby'
      ]
    end

    def supported?(language)
      all.include?(language)
    end
  end 
end
