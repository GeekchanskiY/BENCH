class Article < ApplicationRecord
    validates :title, presence: true,  length: { maximum: 50 }
    validates :body, presence: true,  length: { maximum: 10000 }
end
