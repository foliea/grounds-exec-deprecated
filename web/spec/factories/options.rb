FactoryGirl.define do
  factory :options, class: Hash do
    language 'golang'
    theme 'monokai'
    indent '4'
    keyboard 'vim'

    initialize_with {attributes.stringify_keys}
  end
end
