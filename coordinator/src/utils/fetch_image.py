import aiohttp


async def fetch_image(url: str):
    async with aiohttp.ClientSession() as session:
        async with session.get(url) as response:
            if response.status == 200:
                return await response.read()
            elif response.status == 404:
                raise Exception("Image not found!")
            else:
                raise Exception("Failed to fetch image")
