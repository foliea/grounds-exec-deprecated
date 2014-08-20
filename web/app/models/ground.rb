require 'digest/sha1'

class Ground
  include ActiveModel::Validations
  include ActiveModel::Conversion

  attr_accessor :id, :language, :code

  def initialize(attributes = {})
    attributes.each do |name, value|
      send("#{name}=", value)
    end
  end
  
  def persisted?
    storage.exists(id)
  end
  
  def save
    return if persisted?

    self.id = generate_key
    to_h.each do |field, value|
      storage.hset(id, field, value)
    end
    storage.persist(id)
  end

  def destroy
    storage.del(id)
  end

  def generate_key
    key = 'ground'
    to_h.each do |field, value|
      key << "::#{field}:#{value.to_json}"
    end
    Digest::SHA256.hexdigest(key)
 end

  def self.from_storage(id)
    ground = new(storage.hgetall(id))
    ground.id = id
    ground
  end

  def to_h
    instance_values.slice!('id')
  end

  private

  def self.storage
    $redis
  end
    
  def storage
    self.class.storage
  end
 end
