require 'httparty'

class ObsidianController < ApplicationController
    def index
      @articles = Article.all
      @wasd = "5"
    end
  
    def show
      response = HTTParty.get('http://coordinator_backend/')
      if response.success?
          @additional_argument = "Hello from the controller!"
      else
          @additional_argument = "Failed"
      end
      @article = Article.find(params[:id])
    end
  
    def new
      @article = Article.new
    end
  
    def create
      @article = Article.new(article_params)
      if @article.save
        redirect_to article_path(@article)
      else
        render 'new'
      end
    end
  
    def edit
      @article = Article.find(params[:id])
    end
  
    def update
      @article = Article.find(params[:id])
      if @article.update(article_params)
        redirect_to article_path(@article)
      else
        render 'edit'
      end
    end
  
    def destroy
      @article = Article.find(params[:id])
      @article.destroy
      redirect_to articles_path
    end
  
    private
    def article_params
      params.require(:article).permit(:title, :body)
    end
end 
