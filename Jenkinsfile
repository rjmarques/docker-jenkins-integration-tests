pipeline {
    agent { docker { image 'golang' } }

    environment {
        HOST = 'localhost'
        DB_PASSWORD = 'mysecretpassword'// in real use case go for: https://jenkins.io/doc/book/pipeline/jenkinsfile/#handling-credentials
    }

    stages {
        stage('get go dependencies') {
            steps {
                sh 'go get gotest.tools/assert'
                sh 'go get github.com/lib/pq'
            }
        }
        stage('test') {
            steps {
                sh 'go test'
            }
        }
    }
}