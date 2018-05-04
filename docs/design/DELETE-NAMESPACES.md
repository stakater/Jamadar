# PROBLEM

We would like to delete the dangling/undeed namespaces &/ projects to clean up the cluster.

# SOLUTION

Delete all Namespaces (k8s) & projects (OpenShift) which meet following criteria:

1. Don't contain the annotation `jamadaar.stakater.com/persist=true`
2. Date of creation is older than X period e.g. 1 week (this should be configurable)

Notify on slack when an item is deleted.

This should run regularly and do the cleanup.

Other needs:

- it should work both for vanilla kubernetes & openshift
- it should delete namespaces
- it should delete projects (OpenShift)

So, it will evaluate some expressions and then takes actions; keep in mind this is just first task of Jamadaar and we will be adding a lot more; so, we need to think of doing it in a pluggable way; where one can add new "cleanup" task and Jamadaar should be able to perform it!