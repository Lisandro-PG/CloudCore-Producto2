pipeline {
    agent { label 'cloudcore' }

    stages {
        stage('Construir imagen Docker') {
            steps {
                sh 'docker build -t cloudcore-producto2:latest .'
            }
        }

        stage('Cargar imagen en Minikube') {
            steps {
                sh 'minikube image load cloudcore-producto2:latest'
            }
        }

        stage('Desplegar en Kubernetes') {
            steps {
                sh 'kubectl apply -f k8s/'
                sh 'kubectl rollout restart deployment/cloudcore-app'
                sh 'kubectl rollout status deployment/cloudcore-app'
            }
        }

        stage('Comprobar despliegue') {
            steps {
                sh 'kubectl get pods'
                sh 'kubectl get services'
            }
        }
    }
}
