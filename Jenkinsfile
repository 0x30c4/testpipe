
pipeline {
    agent any

    stages {
        stage('Checkout') {
            steps {
                git 'https://github.com/0x30c4/testpipe'
            }
        }

        stage('Build') {
            steps {
                sh 'go build -o myapp main.go'
            }
        }

        stage('Deploy') {
            steps {
                sh 'nohup ./myapp &'
            }
        }
    }

    post {
        always {
            echo 'Pipeline completed.'
        }
    }
}
