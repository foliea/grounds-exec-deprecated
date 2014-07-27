module ExecCode
  module Language
    extend self

    def all
      [
        'golang',
        'ruby'
      ]
    end

    def supported?(code)
      all.include?(code)
    end
  end 
end
