pipeline {
    agent any

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
            sh 'nohup ./myapp > app.log 2>&1 &'
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
