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
                // Use ssh-agent to manage the SSH connection
                withCredentials([sshUserPrivateKey(credentialsId: 'ppe', keyFileVariable: 'PK')]) {
                    // Stop the existing app on the remote server if running
                    sh '''
                    ssh -i $PK -p 61234 -o StrictHostKeyChecking=no ${REMOTE_USER}@${REMOTE_HOST} << EOF
                        pkill -f myapp || true
                        exit
                    EOF
                    '''

                    // Copy the built binary to the remote server
                    sh 'scp -i $PK -P 61234 -o StrictHostKeyChecking=no myapp ${REMOTE_USER}@${REMOTE_HOST}:${REMOTE_PATH}'

                    // Start the app on the remote server
                    sh '''
                    ssh -i $PK -p 61234 -o StrictHostKeyChecking=no ${REMOTE_USER}@${REMOTE_HOST} << EOF
                        cd ${REMOTE_PATH}
                        nohup ./myapp > app.log 2>&1 &
                        exit
                    EOF
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
