---
sidebar_position: 1
title: Введение в Jenkins
description: Узнайте о Jenkins и его преимуществах.
tags: ["Jenkins", "Continuous Integration", "Continuous Deployment"]
keywords: ["Jenkins", "Continuous Integration", "Continuous Deployment"]
slug: "/jenkins"
---

Jenkins - это сервер автоматизации с открытым исходным кодом. Он помогает автоматизировать части разработки программного обеспечения, связанные со сборкой, тестированием и развертыванием, облегчая непрерывную интеграцию и непрерывную поставку.

## Jenkinsfile - Pipeline как код

Вместо настройки pipeline в Jenkins UI, мы можем определить весь pipeline в Jenkinsfile и зафиксировать его в системе контроля версий. Имя файла чувствительно к регистру и должно называться `Jenkinsfile`.

**Этот pipeline может быть написан двумя способами:**

### Scripted Pipeline

Scripted Pipeline основан на синтаксисе Groovy. Это язык программирования общего назначения с синтаксисом, похожим на Java, но с дополнительными функциями, которые поддерживают общие идиомы программирования.

```Jenkinsfile
node {
    stage('Build') {
        sh 'npm install'
        echo 'Building..'
    }
    stage('Test') {
        echo 'Testing..'
    }
    stage('Deploy') {
        echo 'Deploying..'
    }
}
```

### Declarative Pipeline

Declarative Pipeline - это новый способ определения всего pipeline, используя простую и понятную структуру.

```Jenkinsfile
pipeline {
    agent any 
    stages {
        stage('Build') {
            steps {
                sh npm install
                echo 'Building..'
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying..'
            }
        }
    }
}
```

- **Pipeline**: Директива верхнего уровня, которая определяет весь pipeline.
- **Agent**: Директива agent определяет, где будет выполняться весь pipeline или конкретный этап в среде Jenkins.
- **Stages**: Директива stages определяет различные этапы в pipeline.
- **Stage**: Директива stage определяет один этап в pipeline.
- **Steps**: Директива steps определяет шаги, которые должны быть выполнены в этапе.

Post Step:

Раздел post используется для определения действий, которые должны быть выполнены после завершения pipeline. Он может использоваться для отправки уведомлений, сообщений об успехе/неудаче и т.д.

```Jenkinsfile
stages{}
post{
    always {
        sh 'echo pipeline completed'
    }
    success {
        sh 'echo pipeline sucess'
    }
    failure {
        sh 'echo pipeline failed'
    }
}
```  

When:

Директива when используется для управления потоком pipeline на основе условий. Она может использоваться для пропуска этапов или целых pipeline на основе условий.

```Jenkinsfile
when{
    expression { BRANCH_NAME == 'main' }
}
```

```