module ExecCode
  module Image
    extend self

    def all
      ['exec-golang']
    end

    def image_exist?(name)
      all.include?(name)
    end
  end 
end
