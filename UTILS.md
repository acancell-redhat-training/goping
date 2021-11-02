# QUAY
### Quay login
`podman login quay.io -u "acancell-redhat-training+robot_podman"` # NOTE Use robot *token* as password
# OCP
## OCP PROJECTS
### Create OCP project
`oc new-project "$PROJECT_NAME"`
### Switch to OCP project
`oc project ""$PROJECT_NAME"`
### Destroy OCP project
`oc delete project "$PROJECT_NAME" --force --grace-period=0`
## OCP PODS
### QUICK POD FOR TESTS
`oc run test-bash -it --rm --image=registry.access.redhat.com/ubi8/ubi -- bash`
### USEFUL PACKAGES TO INSTALL ON TEST POD
Package | Tool(s)
------- | -------
`procps-ng` | `free`,`top`
`bind-utils` | `dig`

