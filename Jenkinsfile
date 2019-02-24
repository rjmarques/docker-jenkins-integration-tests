pipeline {
    agent { docker { image 'golang' } }
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