pipeline {
    agent any

    environment {
        IMAGE_NAME = "yourdockerhubuser/golang-gin-app"
        INFRA_REPO = "/tmp/infra"
        GIT_SSH_COMMAND = "ssh -o StrictHostKeyChecking=no"
    }

    stages {
        stage('Checkout') {
            steps {
                git branch: 'main', url: 'git@github.com:hungtd119/notion_clone.git'
            }
        }

        stage('Build Docker Image') {
            steps {
                script {
                    def tag = "v${env.BUILD_NUMBER}"
                    sh "docker build -t $IMAGE_NAME:$tag ."
                    sh "docker push $IMAGE_NAME:$tag"
                    env.NEW_TAG = tag
                }
            }
        }

        stage('Update Infra Repo') {
            steps {
                sshagent(['ssh-credentials']) {
                    sh """
                        rm -rf $INFRA_REPO
                        git clone git@github.com:hungtd119/infra_notion_clone.git $INFRA_REPO
                        cd $INFRA_REPO

                        git checkout main

                        sed -i 's/tag: .*/tag: ${env.NEW_TAG}/' charts/my-app/values.yaml

                        git config user.email "tranduyhungdz119@gmail.com"
                        git config user.name "CI Bot"
                        git add .
                        git commit -m "Update tag to ${env.NEW_TAG}"
                        git push origin main
                    """
                }
            }
        }
    }
}
