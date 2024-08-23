# typed: true
# frozen_string_literal: true

module Foo
  class Bar
    extend T::Sig

    sig { returns(Integer) }
    attr_reader :value

    sig { params(Integer) }
    def initialize(value)
      @value = value
    end

    sig { returns(Integer) }
    def data
      @value
    end
  end
end
