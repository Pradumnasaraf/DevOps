---
sidebar_position: 1
title: Jenkins Introduction
description: Learn about Jenkins and its benefits.
tags: ["Jenkins", "Continuous Integration", "Continuous Deployment"]
keywords: ["Jenkins", "Continuous Integration", "Continuous Deployment"]
slug: "/jenkins"
---

Jenkins is an open-source automation server. It helps automate software delivery tasks such as building, testing, and deploying applications. It is commonly used for CI/CD pipelines.

## Jenkinsfile - Pipeline as Code

Instead of configuring a pipeline only in the Jenkins UI, we can define it in a `Jenkinsfile` and store it in source control. The filename is case-sensitive and should be exactly `Jenkinsfile`.

**This pipeline can be written in two ways:**

### Scripted Pipeline

Scripted Pipeline is based on Groovy syntax. It is a general-purpose programming language with a syntax similar to Java, but with additional features that support common programming idioms.

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

Declarative Pipeline defines the pipeline in a more structured and readable format.

```Jenkinsfile
pipeline {
    agent any 
    stages {
        stage('Build') {
            steps {
                sh 'npm install'
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

- **Pipeline**: The top-level directive that defines the entire pipeline.
- **Agent**: The agent directive defines where the entire pipeline, or a specific stage, will execute in the Jenkins environment.
- **Stages**: The `stages` directive defines the different stages in the pipeline.
- **Stage**: The stage directive defines a single stage in a pipeline.
- **Steps**: The steps directive defines the steps to be executed in a stage.

Post Step:

The `post` section is used to define actions that should happen after the pipeline finishes, such as notifications or success and failure messages.

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

The `when` directive controls whether a stage should run based on conditions.

```Jenkinsfile
when{
    expression { BRANCH_NAME == 'main' }
}
```

```
stage('Building the image') {
    when{
    expression { BRANCH_NAME == 'main' }
    }
    steps {
    sh 'docker build -t pradumnasaraf/devops:latest .'
    }
}
```

### Environment Variables and Credentials

To check default environment variables in Jenkins, we can use the following URL:

```
http://<host>/env-vars.html
http://localhost:8080/env-vars.html
```

### Build Tools

The tools directive is used to define the tools required for the pipeline. It can be used to define the JDK, Maven and Gradle tools.

With the first approach, we can define the tools directly and use their commands in the pipeline.

```jenkinsfile
tools {
    maven 'maven-3.6.3'
    jdk 'jdk-11'
}
```

We can also use wrapper-based steps to define the tools.

```jenkinsfile
stage('Build') {
    steps {
        withMaven(maven: 'maven-3.6.3') {
            sh 'mvn clean install'
        }
    }
}
stage('Build') {
    steps {
        nodejs('node-12.18.3') { // nodejs is the name of the tool
            sh 'npm install'
        }
    }
}
```


Note: The tools must be installed in Jenkins. To install the tools, go to `Manage Jenkins > Global Tool Configuration`. 

### Parameters

The `parameters` directive defines inputs for the pipeline, such as string, boolean, or choice values. In the Jenkins UI, these appear under `Build with Parameters`.

```jenkinsfile
parameters {
    string(name: 'NAME', defaultValue: 'pradumnasaraf', description: 'Enter your name')
    booleanParam(name: 'DEBUG', defaultValue: true, description: 'Enable debug mode')
    choice(name: 'VERSION', choices: ['1.0', '2.0', '3.0'], description: 'Select version')
}
stage('Build') { 
    steps {
        echo "Hello ${params.NAME}" //Accessing the parameters by params.{parameter_name}
        echo "Debug mode is ${params.DEBUG}"
        echo "Version is ${params.VERSION}"
    }
}
```

<img width="1512" alt="Screenshot 2023-01-07 at 1 31 11 PM" src="https://user-images.githubusercontent.com/51878265/211166915-b8e92fad-c5a7-4210-bc20-90b60127e2df.png"></img>


### Triggers

The `triggers` directive is used to define how the pipeline should start. Common triggers include polling and GitHub webhooks.

<img width="1486" alt="Screenshot 2023-01-08 at 12 40 23 AM" src="https://user-images.githubusercontent.com/51878265/211166876-04dfe987-20ae-4a76-8314-27ccf3e636e4.png"></img>



#### Replay

The replay option is used to re-run the pipeline. It is useful when you want to test changes without committing them first.

<img width="1511" alt="Screenshot 2023-01-07 at 2 18 28 PM" src="https://user-images.githubusercontent.com/51878265/211167100-f413eff3-d984-4ef2-a4a5-996b43051d04.png"></img>

### What's next?

- [Learning Resources](./learning-resources.md) - Learn more about Jenkins with these resources.
