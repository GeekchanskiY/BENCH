Rails.application.routes.draw do
  root "obsidian#index"

  resources :articles, path: 'obsidian', controller: 'obsidian'

  get 'users/new'
  get '/users/get_json' => 'obsidian#show_json'
  get '/ping' => 'api#ping'
 
  get "up" => "rails/health#show", as: :rails_health_check
  get "/obsidian", to: "obsidian#index"
  # get "/obsidian/:id", to: "obsidian#show"
 
end
