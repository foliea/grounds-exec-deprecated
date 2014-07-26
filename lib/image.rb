module ExecCode
  module Image
    extend self

    def all
      [
        'exec-golang',
        'exec-ruby'
      ]
    end

    def image_exist?(name)
      all.include?(name)
    end
  end 
end
