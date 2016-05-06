require 'sinatra'
require 'slim'

set :public_folder, 'public'

get '/' do
  slim :index
end
