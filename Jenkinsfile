pipeline {
    agent { docker { image 'golang' } }
    stages {
        stage('test') {
            steps {
                sh 'go test'
            }
        }
    }
}