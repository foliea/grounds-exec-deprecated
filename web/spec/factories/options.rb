FactoryGirl.define do
  factory :options, class: Array do
    [
      ['language', 'golang'],
      ['language', 'python2'],

      ['theme', 'monokai'],
      ['theme', 'github'],

      ['indent', '4'],
      ['indent', 'tab'],

      ['keyboard', 'vim'],
      ['keyboard', 'ace'],
    ]
  end
end
