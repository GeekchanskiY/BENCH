class ObsidianController < ApplicationController
    def index
        @articles = Article.all
        render "index"
    end

    def show
        @article = Article.find(params)
    end
end # TODO: ZERO TRUST ABSOLUTE TRUST
