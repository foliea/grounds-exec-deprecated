class Ground
  include ActiveModel::Validations
  include ActiveModel::Conversion

  KEY_PREFIX = 'ground'

  attr_accessor :id, :language, :code

  def initialize(attributes = {})
    attributes.each do |name, value|
      send("#{name}=", value)
    end
  end

  def persisted?
    id.present?
  end

  def save
    return if persisted?

    self.id = self.class.storage.incr('ground')

    to_h.each do |field, value|
      self.class.storage.hset("#{KEY_PREFIX}:#{id}", field, value)
    end
    puts id
  end

  private

  def self.from_storage(id)
    attributes = storage.hgetall("#{KEY_PREFIX}:#{id}")
    new(attributes)
  end

  def self.storage
    $redis
  end

  def to_h
    instance_values.slice!('id')
  end
 end
