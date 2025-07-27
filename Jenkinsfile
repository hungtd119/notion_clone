pipeline {
    agent any

    environment {
        IMAGE_NAME = "hungit2002/golang-gin-app"
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
                    def imageTag = "${IMAGE_NAME}:${tag}"
                     withCredentials([usernamePassword(credentialsId: 'docker-hub', usernameVariable: 'DOCKER_USER', passwordVariable: 'DOCKER_PASS')]) {
                         sh """
                             echo "\$DOCKER_PASS" | docker login -u "\$DOCKER_USER" --password-stdin
                             docker build -t ${imageTag} .
                             docker push ${imageTag}
                             docker logout
                         """
                     }
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
