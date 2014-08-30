require 'docker'

class ContainersController < ApplicationController
  respond_to :json

  before_action :set_container, only: [:start, :stop, :status]

  def create
    @container = Container.create(params[:language], params[:code])
    render json: { status: :ok , id: @container.id, url: @container.url }
  end

  def start
    @container.start
    render json: { status: :ok }
  end

  def stop
    @container.stop
    render json: { status: :ok }
  end

  def status
    render json: { status: :ok, code: @container.status }
  end

  private

  def set_container
    @container = Container.find_by_id(params['id']) 
  end
end
