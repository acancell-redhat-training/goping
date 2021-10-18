# QUAY
### Quay login
`podman login quay.io -u "acancell-redhat-training+robot_podman"` # NOTE Use robot *token* as password
# OCP
## OCP PROJECTS
### Create OCP project
`oc new-project "acancell-learning"`
### Switch to OCP project
`oc project "acancell-learning"`
### Destroy OCP project
`oc delete project "acancell-learning" --force --grace-period=0`
## OCP PODS
### QUICK POD FOR TESTS
`oc run bashpod -it --rm --image=registry.access.redhat.com/ubi8/ubi -- bash`
