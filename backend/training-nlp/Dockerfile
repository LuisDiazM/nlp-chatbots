FROM python:3.8-slim-buster
WORKDIR /app

COPY . /app
RUN pip install --upgrade pip
RUN pip install -r requirements.txt
RUN python -m nltk.downloader punkt

CMD [ "python", "main.py" ]
