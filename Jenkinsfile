library 'LEAD'
pipeline {
  agent none
  environment {
    SLACK_CHANNEL="cloud-native-demo"
    SLACK_URL = "https://liatrio.slack.com/services/hooks/jenkins-ci/"
  }
  stages {
    stage('Build') {
      agent {
        label "lead-toolchain-skaffold"
      }
      steps {
        notifyPipelineStart()	
        notifyStageStart()
        container('skaffold') {
          sh "skaffold build --file-output=image.json"
          stash includes: 'image.json', name: 'build'
          sh "rm image.json"
        }
      }
    post {	
        success {	
          notifyStageEnd()	
        }	
        failure {	
          notifyStageEnd([result: "fail"])	
        }
      }
    }
    stage('Deploy to Staging') {
      agent {
        label "lead-toolchain-skaffold"
      }
      when {
	      beforeAgent true
	      branch 'main'
      }
      environment {
        NAMESPACE = "${env.stagingNamespace}"
        DOMAIN    = "${env.stagingDomain}"
      }
      steps {
        notifyStageStart()
        container('skaffold') {
          unstash 'build'
          sh "skaffold deploy -a image.json -n ${NAMESPACE}"
        }
        stageMessage "Successfully deployed to staging: `gratibot.${env.stagingDomain}`"
      }
      post {	
        success {	
          notifyStageEnd([status: "Successfully deployed to staging:\ngratibot.${env.stagingDomain}"])	
        }	
        failure {	
          notifyStageEnd([result: "fail"])	
        }	
      }
    }
    stage ('Manual Ready Check') {
      agent none
      when {
        beforeInput true
        branch 'main'
      }
      options {
        timeout(time: 30, unit: 'MINUTES')
      }
      input {
        message 'Deploy to Production?'
      }
      steps {
        echo "Deploying"
      }
    }
    stage('Deploy to Production') {
      when {
        beforeAgent true
        branch 'main'
      }
      agent {
        label "lead-toolchain-skaffold"
      }
      environment {
        NAMESPACE = "${env.productionNamespace}"
        DOMAIN    = "${env.productionDomain}"
      }
      steps {
        notifyStageStart()
        container('skaffold') {
          unstash 'build'
          sh "skaffold deploy -a image.json -n ${NAMESPACE}"
        }
        stageMessage "Successfully deployed to production: `gratibot.${env.productionDomain}`"
      }
    }
  } 		
  post {	
    success {	
      echo "Pipeline Success"	
      notifyPipelineEnd()	
    }	
    failure {	
      echo "Pipeline Fail"	
      notifyPipelineEnd([result: "fail"])	
      }	    
  }
}
