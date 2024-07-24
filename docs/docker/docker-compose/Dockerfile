FROM drupal:9.3.13


RUN apt-get update \
    && apt-get install -y --no-install-recommends \
    git \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /var/www/html/core

COPY ./themes ./themes

WORKDIR /var/www/html
