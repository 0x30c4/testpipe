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
                // Stop the existing app on the remote server if running
                sh '''
                ssh root@123.49.62.187 << EOF
                    pkill -f myapp || true
                    exit
                EOF
                '''

                // Copy the built binary to the remote server
                sh 'scp myapp root@123.49.62.187:/tmp/'

                // Start the app on the remote server
                sh '''
                ssh root@123.49.62.187 << EOF
                    cd /tmp
                    nohup ./myapp > app.log 2>&1 &
                    exit
                EOF
                '''
            }
        }
    }

    post {
        always {
            echo 'Pipeline completed.'
        }
    }
}
