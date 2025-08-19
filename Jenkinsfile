pipeline {
    agent any

    environment {
        PORT = "9090"   // Jenkins will use port 9090
    }

    stages {
        stage('Checkout') {
            steps {
                checkout scm
            }
        }

        stage('Setup Go') {
            steps {
                sh 'go version'
            }
        }

        stage('Build') {
            steps {
                sh 'go mod tidy'
                sh 'go build -o bin/myapp cmd/myapp'
            }
        }

        stage('Run App in Background') {
            steps {
                sh 'nohup ./bin/myapp > app.log 2>&1 & echo $! > app.pid'
            }
        }

        stage('Verify App') {
            steps {
                sh 'go run test_client.go'
            }
        }

    }

    post {
        always {
            echo 'Stopping app if running...'
            sh 'if [ -f app.pid ]; then kill $(cat app.pid) || true; fi'
        }
    }
}
