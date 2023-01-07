def testApp() {
        sh 'echo "Test"'
}
def buildApp() {
        sh 'echo "Build"'
}
def deployApp() {
        sh 'echo "Deploy"'
}

return this
