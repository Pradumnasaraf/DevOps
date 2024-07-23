## Jenkins

Jenkins is an open source automation server. It helps automate the parts of software development related to building, testing, and deploying, facilitating continuous integration and continuous delivery.

### Resources

- [official website](https://www.jenkins.io/)
- [official documentation](https://www.jenkins.io/doc/)
- [Jenkins Docker Image](https://hub.docker.com/r/jenkins/jenkins)
- [Pipeline Syntax](https://www.jenkins.io/doc/book/pipeline/syntax/)

#### Learning Resources

- [Jenkins Docker Installation (docs)](https://www.jenkins.io/doc/book/installing/docker/)
- [Jenkins Tutorial - TechWorld with Nana (video)](https://www.youtube.com/playlist?list=PLy7NrYWoggjw_LIiDK1LXdNN82uYuuuiC)
- [Jenkins FreeCodeCamp (video)](https://youtu.be/f4idgaq2VqA)

## Jenkinsfile - Pipeline as Code

Insted of configuring the pipeline in the Jenkins UI, we can define the entire pipeline in a Jenkinsfile and check it into source control. The file name is case sensitive and must be named `Jenkinsfile`.

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

### Environment Variables and Credentials

To check default environment variables in Jenkins, we can use the following URL:

```
http://<host>/env-vars.html
http://localhost:8080/env-vars.html
```

### Buils Tools

The tools directive is used to define the tools required for the pipeline. It can be used to define the JDK, Maven and Gradle tools.

With the 1st approach we can directly use the tools and its commands in the pipeline.

```jenkinsfile
tools {
    maven 'maven-3.6.3'
    jdk 'jdk-11'
}
```

We can also follow wrapper approach to define the tools.

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

The parameters directive is used to define the parameters required for the pipeline. It can be used to define the parameters like string, boolean, choice, etc. we can check the parameters in the UI by clicking on the `Build with Parameters` button.

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

<img width="1512" alt="Screenshot 2023-01-07 at 1 31 11 PM" src="https://user-images.githubusercontent.com/51878265/211166915-b8e92fad-c5a7-4210-bc20-90b60127e2df.png">


### Triggers

The triggers directive is used to define the triggers for the pipeline. Common way are, poll, github webhooks, etc.

<img width="1486" alt="Screenshot 2023-01-08 at 12 40 23 AM" src="https://user-images.githubusercontent.com/51878265/211166876-04dfe987-20ae-4a76-8314-27ccf3e636e4.png">



#### Replay

The replay option is used to re-run the pipeline. It is useful to test withou making/committing any changes to the code.

<img width="1511" alt="Screenshot 2023-01-07 at 2 18 28 PM" src="https://user-images.githubusercontent.com/51878265/211167100-f413eff3-d984-4ef2-a4a5-996b43051d04.png">

