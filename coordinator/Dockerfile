FROM python:3.10

WORKDIR /app

COPY . .

RUN pip install --no-cache-dir --upgrade -r ./coordinator/requirements.txt

CMD ["uvicorn", "coordinator.src.main:app", "--host", "0.0.0.0", "--port", "80"]
