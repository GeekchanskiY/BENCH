import asyncio
import aiohttp
import logging

from aiohttp.web_response import Response

from classes import Query, Company, ItemPreview, ItemHolder
from classes import get_itemHolder

from selenium import webdriver

from bs4 import BeautifulSoup
from bs4.element import Tag, ResultSet

logger = logging.getLogger(__name__)
logging.basicConfig(encoding='utf-8', level=logging.DEBUG)

async def search_entities(query: str="") -> list:
    return []


async def main():
    q = Query(
        page=1,
        search_query='python',
        # salary=150_000,
        # only_with_salary=True,
        # experience='between1And3',
        area=113
    )

    items: ItemHolder = get_itemHolder()

    async with aiohttp.ClientSession() as session:
        response: Response = await session.get(str(q))

        html: str = await response.text()
        soup = BeautifulSoup(html, 'html.parser')

        current_page: int = 1
        num_pages = int(soup.find('span', attrs={'class': 'pager-item-not-in-short-range'}).text)
        logging.info(f'Pages for query: {num_pages}')
        
        while current_page < num_pages:
            
            q.get_next_page()
            response: Response = await session.get(str(q))

            html: str = await response.text()
            soup: BeautifulSoup = BeautifulSoup(html, 'html.parser')
            page_items: ResultSet = soup.find_all('div', attrs={'class': 'vacancy-serp-item__layout'})

            item: Tag
            for item in page_items:
                new_item_preview = ItemPreview(
                    vacancy_name=item.find('span', attrs={'class': 'serp-item__title'}).text,
                    vacancy_link=item.find('a', attrs={'class': 'bloko-link'})['href'],
                    company_name=item.find('a', attrs={'class': 'bloko-link_kind-tertiary'}).text,
                    company_link=item.find('a', attrs={'class': 'bloko-link_kind-tertiary'})['href'],
                    vacancy_city=item.find('div', attrs={'data-qa': 'vacancy-serp__vacancy-address'}).text,
                    experience_max=0,
                    experience_min=0,
                    remote_job=False,
                    contacts_preview=False
                )
                items.add_item_preview(new_item_preview)

            current_page += 1
        
        logging.info(f'found {len(items)}')
        

if __name__ == '__main__':
    logging.info('Scraper hh.ru start')
    asyncio.run(main())
