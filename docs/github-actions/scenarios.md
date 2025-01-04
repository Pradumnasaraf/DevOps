---
sidebar_position: 2
title: Scenarios and Examples
description: A collection of scenarios and examples to help you get started with GitHub Actions.
tags: ["GitHub Actions", "DevOps", "Continuous Integration", "Continuous Deployment"]
keywords: ["GitHub Actions", "DevOps", "Continuous Integration", "Continuous Deployment"]
slug: "/github-actions/scenarios"
---

This section contains GitHub workflow configuration files for different scenarios. You can use these files as a reference to create your own workflows.

> :warning: Note: The Actions get outdated quickly, so make sure to check the version of the Actions before using them.


### Creating a Release and Tag

This workflow will create a release and a tag when a push is made to the main branch. Below one was used in a Node.js project, but you can use it in any project with a different language.

Key points:
- `version-file` can be changed or can be completely removed if you don't want to use it.

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

### Check and Build Go Project

This workflow will check if the code is getting built and formatted correctly. It will also check for linting issues in the code. Can be useful to check incoming PRs.

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

### Welcome Message

This workflow will send a welcome message to the contributor when they open an issue, PR, or comment on an issue or PR. You can customize the message as per your requirement.

Key points:
- `issue-message` and `pr-message` can be changed as per your requirement.
- `footer` can be changed or can be completely removed if you don't want to use it.

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

This workflow will create a tag and release using GoReleaser when a push is made to the main branch. You can use this workflow in your Go project.

Key points:
- `PA_TOKEN` is a personal access token that you need to create in GitHub with `repo` access and write permissions.
- You should have a `goreleaser.yml` file in the root of your project.

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

### Publish Image to GitHub Container Registry (GHCR)

This workflow will publish the image to the GitHub Container Registry (GHCR) when a release is published. You can use this workflow in your project to publish the image to GHCR.

Key points:
- `tags` can be changed as per your requirement.

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

### Publish Image to Docker Hub

This workflow will publish the image to Docker Hub when a release is published. You can use this workflow in your project to publish the image to Docker Hub.

Key points:
- `DOCKERHUB_USERNAME` and `DOCKERHUB_PASSWORD` are the secrets that you need to create in GitHub with your Docker Hub username and password.

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