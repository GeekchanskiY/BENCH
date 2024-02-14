class ObsidianController < ApplicationController
    def index
        @articles = Article.all
        render json: @articles
    end
end
