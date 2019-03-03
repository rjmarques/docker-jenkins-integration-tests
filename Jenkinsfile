pipeline {
    agent any

    tools {
        go 'Go 1.12'
    }
    
    environment {
        HOST = 'localhost'
        DB_PASS = 'mysecretpassword'// in real use case go for: https://jenkins.io/doc/book/pipeline/jenkinsfile/#handling-credentials
    }

    stages {
        stage('build docker image') {
            steps {
                sh 'docker-compose build'
            }
        }

        stage('run container') {
            steps {
                sh 'docker-compose up -d'
            }
        }

        stage('get go dependencies') {
            steps {
                sh 'go get gotest.tools/assert'
                sh 'go get github.com/lib/pq'
            }
        }

        stage('run tests') {
            steps {
                sh 'go test ./...'
            }
        }
        
        
        stage('stop container and clear images') {
            steps {
                sh 'docker-compose down --rmi all'
            }
        }
    }
}