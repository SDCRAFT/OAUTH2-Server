pipeline {
    agent any

    tools { 
        go 'Go1.23.5' 
        nodejs 'nodeLTS'
    }
    
    stages {
        stage('Checkout') {
            steps {
                git url: 'https://github.com/SDCRAFT/OAuth2-Server.git',
                    branch: 'main'
            }
        }

        stage('Snapshot') {
            steps {
                sh 'curl -sfL https://goreleaser.com/static/run | bash -s -- --snapshot --clean'
            }
        }

        stage('Release') {
            when {
                branch 'master'
                tag pattern: "v*"
            }
            environment {
                GITHUB_TOKEN = credentials('github-token')
            }
            steps {
                sh 'curl -sfL https://goreleaser.com/static/run | bash -s -- --clean'
            }
        }

        stage('Archive') {
            steps {
                archiveArtifacts artifacts: 'dist/Oauth2-Server*', caseSensitive: false 
            }
        }
    }
}