import jenkins.model.*
import hudson.security.*

def instance = Jenkins.getInstance()
def hudsonRealm = new HudsonPrivateSecurityRealm(false)
hudsonRealm.createAccount("admin", "adminPassword")
instance.setSecurityRealm(hudsonRealm)
instance.save()

