module SuperActive
  module Tableless
    def self.included(base)
      base.class_eval do
        include ActiveModel::Validations
        include ActiveModel::Conversion
      end
    end

    def initialize(attributes = {})
      attributes.each do |name, value|
        send("#{name}=", value)
      end
    end

    def persisted?
      false
    end
  end
end
