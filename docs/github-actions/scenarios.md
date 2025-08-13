---
sidebar_position: 2
title: Сценарии и примеры
description: Коллекция сценариев и примеров, чтобы помочь вам начать работу с GitHub Actions.
tags: ["GitHub Actions", "DevOps", "Continuous Integration", "Continuous Deployment"]
keywords: ["GitHub Actions", "DevOps", "Continuous Integration", "Continuous Deployment"]
slug: "/github-actions/scenarios"
---

Этот раздел содержит файлы конфигурации GitHub workflow для различных сценариев. Вы можете использовать эти файлы в качестве справочника для создания собственных рабочих процессов.

> :warning: Примечание: Actions быстро устаревают, поэтому обязательно проверяйте версию Actions перед их использованием.

### Создание релиза и тега

Этот рабочий процесс создаст релиз и тег, когда будет выполнен push в основную ветку. Ниже приведен пример, использованный в проекте Node.js, но вы можете использовать его в любом проекте с другим языком.

Ключевые моменты:
- `version-file` можно изменить или полностью удалить, если вы не хотите его использовать.

```yaml
name: Releases and Changelog
on:
  push:
    branches:
      - main

jobs:
  changelog:
    runs-on: ubuntu-latest

    permissions:
      actions: write
      contents: write

    steps:
      - uses: actions/checkout@v4

      - name: Conventional Changelog Action
        id: changelog
        uses: TriPSs/conventional-changelog-action@v5
        with:
          github-token: ${{ secrets.GITHUB_TOKEN }}
          version-file: './package.json,./package-lock.json'

      - name: Create Release
        if: steps.changelog.outputs.skipped == 'false'
        uses: ncipollo/release-action@v1
        with:
          token: ${{ secrets.GITHUB_TOKEN }}
          tag: ${{ steps.changelog.outputs.tag }}
          name: ${{ steps.changelog.outputs.tag }}
          body: ${{ steps.changelog.outputs.clean_changelog }}
```

### Проверка и сборка Go проекта

Этот рабочий процесс проверит, правильно ли собирается и форматируется код. Он также проверит проблемы с линтингом в коде. Может быть полезен для проверки входящих PR.

```yaml
name: CI

on:
  workflow_dispatch:
  push:
    branches: ["main"]
  pull_request:
    branches: ["main"]

jobs:
  build-format:
    name: Build and Format
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Code
        uses: actions/checkout@v4
      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: 1.22
      - name: Build
        run: go build -v ./...
      - name: Format
        run: diff <(gofmt -d .) <(echo -n)

  golangci-lint:
    name: Lint
    runs-on: ubuntu-latest
    steps:
      - name: Checkout Code
        uses: actions/checkout@v4
      - name: Setup Go
        uses: actions/setup-go@v5
        with:
          go-version: 1.22
          cache: false
      - name: GolangCI-Lint Check
        uses: golangci/golangci-lint-action@v6
```

### Приветственное сообщение

Этот рабочий процесс отправит приветственное сообщение участникам, когда они открывают проблему, PR или комментируют проблему или PR. Вы можете настроить сообщение по вашему усмотрению.

Ключевые моменты:
- `issue-message` и `pr-message` можно изменить по вашему усмотрению.
- `footer` можно изменить или полностью удалить, если вы не хотите его использовать.

```yaml
name : Greetings

on:
  fork:
  push:
    branches: [main]
  issues:
    types: [opened]
  issue_comment:
    types: [created]
  pull_request_target:
    types: [opened]
  pull_request_review_comment:
    types: [created]

jobs:
  welcome:
    name: Welcome Step
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: EddieHubCommunity/gh-action-community/src/welcome@main
        with:
          github-token: ${{ secrets.github_token }}
          issue-message: "<h3>Hey! contributor, thank you for opening an Issue 🎉.</h3>"
          pr-message: "<h3>Hey! contributor, thank you for opening a Pull Request 🎉.</h3>"
          footer: "<h3>Thank you for contributing to the project 🚀</h3>"
```

### GoReleaser

Этот рабочий процесс создаст тег и релиз с помощью GoReleaser, когда будет выполнен push в основную ветку. Вы можете использовать этот рабочий процесс в вашем Go-проекте.

Ключевые моменты:
- `PA_TOKEN` - это личный токен доступа, который вы должны создать в GitHub с правами доступа `repo` и разрешениями на запись.
- У вас должен быть файл `goreleaser.yml` в корне вашего проекта.

```yaml
name: GoReleaser
on:
  push:
    branches:
      - main

jobs:
  create-tag:
    name: Creating tag
    runs-on: ubuntu-latest
    steps:
      - name: Checkout the repo
        uses: actions/checkout@v4

      - name: Creating Changelog and Release
        id: changelog
        uses: TriPSs/conventional-changelog-action@v5
        with:
          github-token: ${{ secrets.PA_TOKEN }}
          output-file: "false"
          skip-commit: "true"

    outputs:
      tag: ${{ steps.changelog.outputs.tag }}

  goreleasers:
    name: GoReleaser
    runs-on: ubuntu-latest
    needs: create-tag
    # only run if the tag is not empty
    if: ${{ needs.create-tag.outputs.tag != '' }}
    steps:

      - name: Checkout the repo
        uses: actions/checkout@v4
        with:
          fetch-depth: 0

      - name: Set up Go
        uses: actions/setup-go@v5

      - name: Run GoReleaser
        uses: goreleaser/goreleaser-action@v4
        with:
          distribution: goreleaser
          version: latest
          args: release --clean
        env:
          GITHUB_TOKEN: ${{ secrets.GITHUB_TOKEN }}
```

### Публикация изображения в GitHub Container Registry (GHCR)

Этот рабочий процесс опубликует изображение в GitHub Container Registry (GHCR), когда будет опубликован релиз. Вы можете использовать этот рабочий процесс для публикации изображения в GHCR.

Ключевые моменты:
- `tags` можно изменить по вашему усмотрению.

```yaml
name: Publish Image - GHCR

on:
  workflow_dispatch:
  release:
    types: [published]

env:
  REGISTRY: ghcr.io
  IMAGE_NAME: ${{ github.repository }}

jobs:
  build-and-push-image:
    runs-on: ubuntu-latest
    permissions:
      contents: read
      packages: write

    steps:
      - name: Checkout repository
        uses: actions/checkout@v4

      - name: get-npm-version
        id: package-version
        uses: martinbeentjes/npm-get-version-action@v1.3.1

      - name: Log in to the Container registry
        uses: docker/login-action@v3
        with:
          registry: ${{ env.REGISTRY }}
          username: ${{ github.actor }}
          password: ${{ secrets.GITHUB_TOKEN }}

      - name: Extract metadata (tags, labels) for Docker
        id: meta
        uses: docker/metadata-action@v5
        with:
          images: ${{ env.REGISTRY }}/${{ env.IMAGE_NAME }}

      - name: Build and push Docker image
        uses: docker/build-push-action@v6
        with:
          context: .
          push: true
        # This is format ghcr.io/{owner}/{repo}:{tag} and ghcr.io/{owner}/{repo}:latest, change it as per your requirement
          tags: | 
            ghcr.io/pradumnasaraf/devops:${{ steps.package-version.outputs.current-version}}
            ghcr.io/pradumnasaraf/devops:latest
          labels: ${{ steps.meta.outputs.labels }}
```

### Публикация изображения в Docker Hub

Этот рабочий процесс опубликует изображение в Docker Hub, когда будет опубликован релиз. Вы можете использовать этот рабочий процесс для публикации изображения в Docker Hub.

Ключевые моменты:
- `DOCKERHUB_USERNAME` и `DOCKERHUB_PASSWORD` - это секреты, которые вы должны создать в GitHub с вашим именем пользователя Docker Hub и паролем.

```yaml
name: Publish Image to DockerHub

on:
  workflow_dispatch:
  release:
    types: [published]

jobs:
  build:
    runs-on: ubuntu-latest

    steps:
      - uses: actions/checkout@v4

      - name: get-npm-version
        id: package-version
        uses: martinbeentjes/npm-get-version-action@v1.3.1

      - name: DockerHub Login
        uses: docker/login-action@v3
        with:
          username: ${{ secrets.DOCKERHUB_USERNAME }}
          password: ${{ secrets.DOCKERHUB_PASSWORD }}

      - name: Build the Docker image
        run: docker build . --file Dockerfile --tag ${{ secrets.DOCKERHUB_USERNAME }}/devops:${{ steps.package-version.outputs.current-version}}

      - name: Docker Push
        run: docker push ${{ secrets.DOCKERHUB_USERNAME }}/devops:${{ steps.package-version.outputs.current-version}}
```