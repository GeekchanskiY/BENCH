import asyncio
import aiohttp
import logging

from classes import Query, Company, ItemPreview

from selenium import webdriver

logger = logging.getLogger(__name__)
logging.basicConfig(encoding='utf-8', level=logging.DEBUG)

async def search_entities(query: str="") -> list:
    return []


async def main():
     async with aiohttp.ClientSession() as session:
        async with session.get('https://hh.ru/search/vacancy?text=python&area=2814&hhtmFrom=main&hhtmFromLabel=vacancy_search_line') as response:

            print("Status:", response.status)
            print("Content-type:", response.headers['content-type'])

            html = await response.text()
            with open('out.txt', 'w', encoding='utf-8') as f:
                f.write(html)
            # print(html)

if __name__ == '__main__':
    logging.info('Scraper hh.ru start')
    asyncio.run(main())
