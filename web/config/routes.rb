Rails.application.routes.draw do
  root 'grounds#show'

  match 's/:id', to: 'grounds#shared', as: 'grounds_shared', via: :get
  match 'grounds/share/', to: 'grounds#share', as: 'grounds_share', via: :post
  match 'grounds/switch_option/:option', to: 'grounds#switch_option', as: 'grounds_switch_option', via: :put
end
