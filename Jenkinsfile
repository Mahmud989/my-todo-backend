pipeline {
    agent any

    stages {

        stage ('Git - Checkout') {
            steps {
                checkout([$class: 'GitSCM', branches: [[name: '*/master']], doGenerateSubmoduleConfigurations: false,
                 extensions: [], submoduleCfg: [], userRemoteConfigs: [[credentialsId: 'Jenkins-git-auth',
                  url: 'https://github.com/Mahmud989/go-signalling-server.git']]])
            }
        }

        stage ('Docker build') {
            steps {
                sh '''
                    docker build -t hub.letsecure.az/go-signalling-server .
                '''
            }
        }

        stage ('Docker push') {
            steps {
                sh 'docker push hub.letsecure.az/go-signalling-server'
            }
        }

        stage ('Docker rm container') {
            steps {
                catchError(buildResult: 'SUCCESS', stageResult: 'FAILURE') {
                    sh "exit 0"
                }
                sh 'docker-compose rm -f -s go-signalling-server  && echo "container go-general-scrapper removed" || echo "container go-general-scrapper does not exist"'
            }
        }

        stage ('Docker up') {
            steps {
                sh 'docker-compose up -d go-signalling-server'
            }
        }
    }
}