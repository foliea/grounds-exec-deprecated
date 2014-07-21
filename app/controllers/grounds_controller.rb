class GroundsController < ApplicationController
  def show
    @ground = Ground.new
  end

  def run
    binding.pry     
  end
end
