FactoryGirl.define do
  factory :options, class: Hash do
    language 'golang'
    theme 'monokai'
    indent 'tab'

    initialize_with {attributes.stringify_keys}
  end
end
