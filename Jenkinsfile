pipeline {
    agent any   // runs on any available agent (or master if allowed)

    stages {
        stage('Verify Go') {
            steps {
                bat 'where go'
                bat 'go version'
            }
        }

        stage('Build') {
            steps {
                bat 'go build -o bin\\myapp.exe cmd\\myapp'
            }
        }
    }
}
