class ObsidianController < ApplicationController
    def index
        @articles = Article.all
    end
end
