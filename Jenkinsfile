#!groovy
@Library("platform.infrastructure.pipelinelibrary") _

utils = new com.ebsco.platform.Utils()
artifactory = new com.ebsco.platform.datapipeline.ArtifactoryHelper()

nodeId = platformDefaults.getBuildNodeLabel()
credentialsId = platformDefaults.getCredentialsId()

node(nodeId) {

  step([$class: 'WsCleanup'])//delete project workspace
    // Checkout this repo from github
    stage ("checkout") {
        checkout scm
    }

    stage('Build Test Container') {
        docker.build("infoblox-test-container", "-f Dockerfile . -o .")
        sh "ls -la"
    }

    // if(env.BRANCH_NAME == "main" || env.BRANCH_NAME.startsWith("PR-")) {
    //     stage('Publish to artifactory') {
    //         artifactory.publishArchive("terraform-provider-infoblox", )
    //     }
    // }
}