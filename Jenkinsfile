pipeline {
    agent any

    environment {
      REMOTE_USER = 'root'  // Replace with your remote server's username
      REMOTE_HOST = '123.49.62.187'  // Replace with your remote server's IP
      REMOTE_PATH = '/tmp/'  // Replace with your deployment directory on the remote server
    }

    stages {
        stage('Checkout') {
          steps {
            git branch: 'main', url: 'https://github.com/0x30c4/testpipe'
          }
        }

        stage('Build') {
          steps {
            sh 'go build -o myapp main.go'
          }
        }

        stage('Deploy') {
          steps {

          }
        }
      }
    }

    post {
        always {
            echo 'Pipeline completed.'
        }
    }
}
