# QUAY
### Quay login
`podman login **quay.io** -u "acancell-redhat-training+robot_podman"` # NOTE Use robot *token* as password
# OCP
### Create OCP project
`oc new-project "acancell-learning"`
### Switch to OCP project
`oc project "acancell-learning"`
### Destroy OCP project
`oc delete project "acancell-learning" --force --grace-period=0`
