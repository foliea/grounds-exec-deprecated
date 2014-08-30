Rails.application.routes.draw do
  root 'grounds#show'

  match 'grounds', to: 'grounds#show', as: 'ground_show', via: :get
  match 's/:id', to: 'grounds#shared', as: 'ground_shared', via: :get 
  match 'grounds/share', to: 'grounds#share', as: 'ground_share', via: :post
  match 'grounds/switch_option', to: 'grounds#switch_option', as: 'ground_switch_option', via: :put

  match 'containers/create', to: 'containers#create', as: 'container_create', via: :post
  match 'containers/start', to: 'containers#start', as: 'container_start', via: :post
  match 'containers/stop', to: 'containers#stop', as: 'container_stop', via: :post
  match 'containers/remove', to: 'containers#remove', as: 'container_remove', via: :post
  match 'containers/status', to: 'containers#status', as: 'container_status', via: :get
end
