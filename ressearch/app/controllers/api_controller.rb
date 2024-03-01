class ApiController < ApplicationController
    def ping
        render json: {"msg": "pong"}
    end
end
