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
            // SSH into the remote server and deploy the application
            sshCommand remote: [
              user: REMOTE_USER,
              host: REMOTE_HOST,
              credentialsId: 'ppe',
              trustUnknownHosts: true
            ], command: '''
              pkill -f myapp || true
              scp myapp ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_PATH}
              cd ${REMOTE_PATH}
              nohup ./myapp > app.log 2>&1 &
            '''
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
