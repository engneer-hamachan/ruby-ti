# frozen_string_literal: true

# Plan Value Object
class Plan
  attr_reader :name, :content

  def initialize(name:, content:)
    @name = name
    @content = content
  end

  def self.new(name:, content: nil)
    return nil, 'name is must be a string' unless name.is_a?(String)
    return nil, 'name minimum length is 1' if name.empty?

    super(name: name, content: content)
  end
end

a = Plan.new(name: '1', content: '2')
dbtp a.name
