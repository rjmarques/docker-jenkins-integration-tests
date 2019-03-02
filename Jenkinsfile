pipeline {
    agent any

    stages {
        stage('build docker image') {
            steps {
                sh 'docker-compose build'
            }
        }
        
        stage('run container') {
            steps {
                sh 'docker-compose up'
            }
        }

        stage('run container') {
            steps {
                sh 'docker-compose up'
            }
        }
        
        stage('stop container and clear images') {
            steps {
                sh 'docker-compose down --rmi all'
            }
        }
    }
}