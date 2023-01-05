## Jenkins

Jenkins is an open source automation server. It helps automate the parts of software development related to building, testing, and deploying, facilitating continuous integration and continuous delivery.

### Resources

- [official website](https://www.jenkins.io/)
- [official documentation](https://www.jenkins.io/doc/)
- [Jenkins Docker Image](https://hub.docker.com/r/jenkins/jenkins)
- [Pipeline Syntax](https://www.jenkins.io/doc/book/pipeline/syntax/)

##### Learning Resources

- [Jenkins Docker Installation (docs)](https://www.jenkins.io/doc/book/installing/docker/)
- [Jenkins Tutorial - TechWorld with Nana (video)](https://www.youtube.com/playlist?list=PLy7NrYWoggjw_LIiDK1LXdNN82uYuuuiC)
- [Jenkins FreeCodeCamp (video)](https://youtu.be/f4idgaq2VqA)
- [Jenkins Tutorial - Guru99 (video)](https://youtu.be/5XQOK0v_YRE)

### Jenkisfile - Pipeline as Code

Insted of configuring the pipeline in the Jenkins UI, we can define the entire pipeline in a Jenkinsfile and check it into source control. The the file name is case sensitive and must be named `Jenkinsfile`.

**This pipeline can be written in two ways:**

#### Scripted Pipeline

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

#### Declarative Pipeline

Declarative Pipeline is a new way of defining the entire pipeline using a simple and easy to understand structure.

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

- **Pipeline**: The top-level directive that defines the entire pipeline.
- **Agent**: The agent directive defines where the entire pipeline, or a specific stage, will execute in the Jenkins environment.
- **Stages**: The stages directive defines the different stages in pipeline.
- **Stage**: The stage directive defines a single stage in a pipeline.
- **Steps**: The steps directive defines the steps to be executed in a stage.

Post Step:

The post section is used to define actions to be taken after the completion of the pipeline. It can be used to send notifications, sccess/failure messages, etc.

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

The when directive is used to control the flow of the pipeline based on the conditions. It can be used to skip stages, or entire pipelines, based on the conditions.

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

### Environment Variables

To check default environment variables in Jenkins, we can use the following URL:

```
http://<host>/env-vars.html
http://localhost:8080/env-vars.html
```








