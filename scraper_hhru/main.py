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
        # salary=150_000,
        # only_with_salary=True,
        # experience='between1And3',
        area=113
    )
     print(str(q))
     async with aiohttp.ClientSession() as session:
        response = await session.get(str(q))
        print(str(q))
        print("Status:", response.status)
        print("Content-type:", response.headers['content-type'])

        html = await response.text()
        soup = BeautifulSoup(html, 'html.parser')
        print(soup.find('div', attrs={'class': 'vacancysearch-xs-header'}).text)
        print(soup.find_all('div', attrs={'class': 'vacancy-serp-item__layout'}).__len__())
        current_page = 1
        links: list = list()
        num_pages = int(soup.find('span', attrs={'class': 'pager-item-not-in-short-range'}).text)
        print(num_pages)
        
        while current_page < 4:
            
            q.get_next_page()
            print(str(q))
            response = await session.get(str(q))


            html = await response.text()
            soup = BeautifulSoup(html, 'html.parser')
            # print(soup.find('div', attrs={'class': 'vacancysearch-xs-header'}).text)
            items = soup.find_all('div', attrs={'class': 'vacancy-serp-item__layout'})
            for item in items:
                links.append(item.find('a', attrs={'class': 'bloko-link'})['href'])
                # print(item.find('a', attrs={'class': 'bloko-link'}).text)
            current_page += 1
        links = list(set(links))
        print(links)
        print(len(links))
        


            
            # with open('out.txt', 'w', encoding='utf-8') as f:
            #     f.write(html)
            # print(html)

if __name__ == '__main__':
    logging.info('Scraper hh.ru start')
    asyncio.run(main())
