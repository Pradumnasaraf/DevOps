## Jenkins

Jenkins is an open source automation server. It helps automate the parts of software development related to building, testing, and deploying, facilitating continuous integration and continuous delivery.

### Resources

- [official website](https://www.jenkins.io/)
- [official documentation](https://www.jenkins.io/doc/)
- [Docker Image](https://hub.docker.com/r/jenkins/jenkins)
- [Octopus Docker Installation](https://octopus.com/blog/jenkins-docker-install-guide)

### Jenkisfile - Pipeline as Code

Jenkinsfile is a file that contains the definition of a Jenkins Pipeline and is checked into source control. Insted of configuring the pipeline in the Jenkins UI, you can define the entire pipeline in a Jenkinsfile and check it into source control.

The the file name is case sensitive and must be named `Jenkinsfile`.

Syntax:

```groovy

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

- Pipeline: The top-level directive that defines the entire pipeline.
- Agent: The agent directive defines where the entire pipeline, or a specific stage, will execute in the Jenkins environment.
- Stages: The stages directive defines the different stages in pipeline.
- Stage: The stage directive defines a single stage in a pipeline.
- Steps: The steps directive defines the steps to be executed in a stage.
