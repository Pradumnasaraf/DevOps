---
sidebar_position: 1
title: Введение в GitHub Actions
description: Руководство по изучению GitHub Actions.
tags: ["GitHub Actions", "DevOps", "CI/CD"]
keywords: ["GitHub Actions", "DevOps", "CI/CD"]
slug: "/github-actions"
---

GitHub Actions - это функция, которая позволяет автоматизировать ваши рабочие процессы разработки программного обеспечения. Вы можете писать отдельные задачи, называемые действиями, и объединять их для создания пользовательского рабочего процесса. Рабочие процессы - это пользовательские автоматизированные процессы, которые вы можете настроить в своем репозитории для сборки, тестирования, упаковки, выпуска или развертывания любого проекта кода на GitHub.

### Обзор

- [Документация](https://docs.github.com/en/actions/learn-github-actions/understanding-github-actions)

- **Workflow (Рабочий процесс)** - Рабочий процесс - это настраиваемый автоматизированный процесс, состоящий из одной или нескольких задач. Рабочие процессы определяются в файлах `.yml` в директории `.github/workflows` вашего репозитория.

- **Jobs (Задачи)** - Задача - это набор шагов, которые выполняются на одном и том же runner. Runner - это сервер, на котором установлено приложение GitHub Actions runner.

- **Steps (Шаги)** - Шаг - это отдельная задача, которая может запускать команды или действия, такие как `actions/checkout@v2`. Каждый шаг в задаче выполняется на одном и том же runner, что позволяет напрямую обмениваться файлами.

> Резюме: Рабочий процесс - это набор задач, и каждая задача - это набор шагов. Каждый шаг может быть действием или shell командой.

Базовый файл рабочего процесса.

```yaml
name: CI # имя рабочего процесса

on: [push] # запускает рабочий процесс при push или pull request событиях, но только для ветки master

jobs:
  build: # имя задачи
    runs-on: ubuntu-latest # runs-on - это тип машины для запуска задачи - runner
    steps: # шаги - это отдельные задачи, которые составляют задачу
      - uses: actions/checkout@v2
      - name: Run a one-line script # name - это имя шага
        run: echo Hello, world!
```

- **Action (Действие)** - Действие - это пользовательское приложение для платформы GitHub Actions, которое выполняет сложную, но часто повторяющуюся задачу.

- **Event (Событие)** - Событие - это конкретная активность, которая запускает рабочий процесс. Например, активность, которая происходит на GitHub, такая как открытие pull request или отправка коммита.

### Триггеры

Ключевое слово `on` используется для запуска рабочего процесса. Он может быть запущен push, pull request или расписанием.

- [Документация](https://docs.github.com/en/actions/reference/events-that-trigger-workflows)

```yaml
on:
  push:
    branches: [main]
  pull_request:
    branches: [main]
```

#### Фильтры триггеров

- Мы можем запускать рабочий процесс только когда изменяются определенные файлы.

```yaml
on:
  push:
    paths:
      - "src/index.js"
```

- Мы можем запускать рабочий процесс только когда изменяются определенные файлы с расширением.

```yaml
on:
  push:
    paths:
      - "**.js"
      - "**.css"
  pull_request:
    paths:
      - "**.js"
```

- Мы можем запускать рабочий процесс только когда изменяются определенные файлы в директории.

```yaml
on:
  push:
    paths:
      - "src/**"
```

- Ручной запуск

```yaml
on:
  workflow_dispatch:
```

### Сервисы

Сервисы - это внешние ресурсы, которые можно использовать в рабочем процессе. Мы можем использовать сервисы, такие как `redis` или `postgres` в нашем рабочем процессе. Мы можем запустить сервис в контейнере и использовать его в нашем рабочем процессе. Мы также можем использовать пользовательские образы.

```yaml
services:
  redis:
    image: redis # Docker image
    ports:
      - 6379:6379 # port mapping

  mongodb:
    image: mongo
    ports:
      - 27017:27017

  sqlite:
    image: sqlite
    ports:
      - 3306:3306
```

### Использование вывода одного шага в другом шаге

Мы можем использовать вывод из одного шага в другой шаг, используя ключевое слово `id`.

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Get the version
        id: get_version
        run: echo ::set-output name=version::1.0.0

      - name: Use the version
        run: echo ${{ steps.get_version.outputs.version }} # {{ steps.<step_id>.outputs.<output_name> }}
```

### Задачи

В рабочем процессе можно иметь несколько задач. Каждая задача запускается параллельно по умолчанию. Мы также можем запускать задачи последовательно, используя ключевое слово `needs`. Каждый шаг в задаче выполняется на одном и том же runner.

Мы можем запускать задачу, когда задача проходит или завершается с ошибкой, используя ключевое слово `needs`.

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Get the version
        run: echo "Hello World"

  test:
    needs: build
    runs-on: ubuntu-latest
    steps:
      - name: Get the version
        run: echo "Hello World"
```

### GitHub context

Контекст `github` доступен вам в любом рабочем процессе или действии, которое вы создаете на GitHub. Вы можете использовать контекст для получения информации о рабочем процессе, репозитории, событии, которое запустило рабочий процесс, например, номер pull request и т.д.

**Некоторые полезные свойства контекста `github` являются:**

- `github.event_name` - Имя вебхука события, которое запустило рабочий процесс.
- `github.sha` - SHA коммита, который запустил рабочий процесс.
- `github.ref` - Ссылка на ветку или тег, который запустил рабочий процесс.
- `github.repository` - Владелец и имя репозитория. Например, `Pradumnasaraf/DevOps`.
- `github.actor` - Имя человека или приложения, которое инициировало рабочий процесс.
- `github.job` - Имя задачи, которая в данный момент выполняется.
- `github.run_number` - Уникальный номер для каждого запуска определенного рабочего процесса в репозитории. Этот номер начинается с 1 для первого запуска рабочего процесса, и увеличивается с каждым новым запуском. Этот номер не изменяется, если вы повторно запускаете запуск рабочего процесса.
- `github.run_id` - Уникальный номер для каждого запуска любого рабочего процесса в репозитории. Этот номер начинается с 1 для первого запуска рабочего процесса, и увеличивается с каждым новым запуском. Этот номер не изменяется, если вы повторно запускаете запуск рабочего процесса.
- `github.workflow` - Имя рабочего процесса.
- `github.action` - Уникальный идентификатор (id) действия.
- `github.event` - Загрузка события. Например, объект issue или pull request, который запустил рабочий процесс.

### Переменные окружения

Мы можем устанавливать переменные окружения в файле рабочего процесса, используя ключевое слово `env`. WE

```yaml
env:
  MY_NAME: "Pradumna"
```

### Секреты

Мы можем хранить чувствительные данные, такие как API-ключи, пароли и т.д., в настройках репозитория. Мы можем получить доступ к секретам, используя контекст `secrets`.

```yaml
jobs:
  build:
    runs-on: ubuntu-latest
    steps:
      - name: Get the version
        run: echo ${{ secrets.MY_SECRET }}
```

### Матрица

Матрица полезна, когда мы хотим запустить задачу на нескольких версиях инструмента.Например, мы хотим запустить задачу на нескольких версиях Node.js. Мы можем использовать ключевое слово `matrix` для запуска задачи на нескольких версиях инструмента. Мы также можем использовать ключевое слово `matrix` для запуска задачи на нескольких операционных системах.

```yaml
jobs:
  example_matrix:
    strategy:
      matrix:
        os: [ubuntu-22.04, ubuntu-20.04]
        version: [14, 16, 18]
    runs-on: ${{ matrix.os }}
    steps:
      - uses: actions/setup-node@v3
        with:
          node-version: ${{ matrix.version }}
```

### Выходы

Мы можем использовать ключевое слово `outputs` для вывода данных из задачи. Мы можем использовать вывод из одной задачи в другую задачу, используя ключевое слово `needs`.

```yaml
jobs:
  deploy:
    runs-on: ubuntu-latest
    outputs:
      url: ${{ steps.deploy-preview.outputs.url }}
    steps:
      - uses: actions/checkout@v3
      - id: deploy-preview
        run: echo "url=preview_url" >> $GITHUB_OUTPUT

    data:
      runs-on: ubuntu-latest
      needs: deploy
      steps:
        - name: load json files
          run: echo ${{ needs.deploy.outputs.url }}
```

## Артефакты

Он позволяет обмениваться данными между задачами в рабочем процессе и хранить данные после завершения рабочего процесса. Мы можем использовать действия `upload-artifact` и `download-artifact` для загрузки и загрузки артефактов.

```yaml
name: Share data between jobs

on: [push]

jobs:
  job_1:
    name: Generate Random Number
    runs-on: ubuntu-latest
    steps:
      - shell: bash
        run: |
          echo $((RANDOM % 100)) > random-number.txt
      - name: Upload random number for job 1
        uses: actions/upload-artifact@v4
        with:
          name: random_number
          path: random-number.txt

  job_2:
    name: Echo Random Number
    needs: job_1
    runs-on: ubuntu-latest
    steps:
      - name: Download random number from job 1
        uses: actions/download-artifact@v4
        with:
          name: random_number
      - shell: bash
        run: |
          number=$(cat random-number.txt)
          echo The random number is $number
``` 


### Что дальше?

- [Сценарии](./scenarios.md) - Коллекция файлов рабочих процессов GitHub Actions, которые я использую и создал, чтобы помочь вам лучше понять концепции.
- [Обучающие ресурсы](./learning-resources.md) - Список ресурсов для изучения GitHub Actions.
- [Другие ресурсы](./other-resources.md) - Список других ресурсов, которые вы можете использовать для изучения GitHub Actions.