pipeline {
    agent any

    stages {

        stage ('Git - Checkout') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false,
                 extensions: [], submoduleCfg: [], userRemoteConfigs: [[credentialsId: 'Jenkins-git-auth',
                  url: 'https://github.com/Mahmud989/my-todo-backend.git']]])
            }
        }

        stage ('Docker build') {
            steps {
                sh '''
                    docker build -t hub.letsecure.az/my-todo-backend .
                '''
            }
        }

        stage ('Docker push') {
            steps {
                sh 'docker push hub.letsecure.az/my-todo-backend'
            }
        }

        stage ('Docker rm container') {
            steps {
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    sh "exit 0"
                }
                sh 'docker-compose rm -f -s my-todo-backend  && echo "container go-general-scrapper removed" || echo "container my-todo-backend does not exist"'
            }
        }

        stage ('Docker up') {
            steps {
                sh 'docker-compose up -d my-todo-backend'
            }
        }
    }
}