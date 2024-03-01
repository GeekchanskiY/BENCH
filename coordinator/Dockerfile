FROM python:3.10

WORKDIR /app

COPY . .

RUN pip install -r ./requirements.txt

WORKDIR /app/src
CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "80", "--reload"]
