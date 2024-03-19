import asyncio
import aiohttp
import logging

from classes import Query, Company, ItemPreview

from selenium import webdriver

from bs4 import BeautifulSoup

logger = logging.getLogger(__name__)
logging.basicConfig(encoding='utf-8', level=logging.DEBUG)

async def search_entities(query: str="") -> list:
    return []


async def main():
     q = Query(
        page=1,
        search_query='python',
        salary=150_000,
        only_with_salary=True,
        experience='between1And3',
        area=113
    )
     async with aiohttp.ClientSession() as session:
        async with session.get(str(q)) as response:

            print("Status:", response.status)
            print("Content-type:", response.headers['content-type'])

            html = await response.text()
            soup = BeautifulSoup(html, 'html.parser')
            print(soup.find_all('span', attrs={'class': 'serp-item__title'}))
            # with open('out.txt', 'w', encoding='utf-8') as f:
            #     f.write(html)
            # print(html)

if __name__ == '__main__':
    logging.info('Scraper hh.ru start')
    asyncio.run(main())
