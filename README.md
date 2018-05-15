# Jamadaar

[![Go Report Card](https://goreportcard.com/badge/github.com/stakater/jamadaar?style=flat-square)](https://goreportcard.com/report/github.com/stakater/jamadaar)
[![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/stakater/jamadaar)
[![Release](https://img.shields.io/github/release/stakater/jamadaar.svg?style=flat-square)](https://github.com/stakater/jamadaar/releases/latest)
[![GitHub tag](https://img.shields.io/github/tag/stakater/jamadaar.svg?style=flat-square)](https://github.com/stakater/jamadaar/releases/latest)
[![Docker Pulls](https://img.shields.io/docker/pulls/stakater/jamadaar.svg?style=flat-square)](https://hub.docker.com/r/stakater/jamadaar/)
[![Docker Stars](https://img.shields.io/docker/stars/stakater/jamadaar.svg?style=flat-square)](https://hub.docker.com/r/stakater/jamadaar/)
[![MicroBadger Size](https://img.shields.io/microbadger/image-size/stakater/jamadaar.svg?style=flat-square)](https://microbadger.com/images/stakater/jamadaar)
[![MicroBadger Layers](https://img.shields.io/microbadger/layers/stakater/jamadaar.svg?style=flat-square)](https://microbadger.com/images/stakater/jamadaar)
[![license](https://img.shields.io/github/license/stakater/jamadaar.svg?style=flat-square)](LICENSE)

[![Get started with Stakater](https://stakater.github.io/README/stakater-github-banner.png)](http://stakater.com/?utm_source=Jamadaar&utm_medium=github)


## WHY NAME JAMADAAR?
Jamadaar, an Urdu word, is used for Sweepers/Cleaners in Pakistan. This Jamadaar will keep your cluster clean and sweep away the left overs of your cluster and will act as you want it to.

## Problem
Dangling/Redundant resources take a lot of space and memory in a cluster. So we want to delete these unneeded resources depending upon the age and pre-defined  annotations. e.g. I would like to delete namespaces that were without a specific annotation and are almost a month old and would like to take action whenever that happens.

## Solution

Jamadaar is a Kubernetes controller that can poll at configured time intervals and watch for dangling resources that are an 'X' time period old and don't have a specific annotation, and will delete them and take corresponding actions.

## Configuring

First of all you need to modify `configs/config.yaml` file. Following are the available options that you can use to customize Jamadaar:

| Key                   |Description                                                                    |
|-----------------------|-------------------------------------------------------------------------------|
| polltimeinterval      | The time interval after which the controller will poll and look for dangling resources, The value can be in "ms", "s", "m", "h" or even combined like 2h45m       |
| age        | The time period that a dangling resource  has been created e.g. delete only resources that are 7 days old, The value can be in "d", "w", "m", "y", Combined format is not supported     |
| actions               | The Array of actions that you want to take, e.g. send message to Slack, etc   |

### Supported Resources
Currently we are supporting the following dangling resources,
- Namespaces


We will be adding support for other Resources as well in the future

### Supported Actions
Currently we are supporting following Actions with their Parameters,
- Default: No parameters needed, it will just log to console the details.
- Slack: you need to provide `token` and `Channel Name` as Parameters in the yaml file

We will be adding support for other Actions as well in the future

## Deploying to Kubernetes

You have to first clone or download the repository contents. The kubernetes deployment and files are provided inside `deployments/kubernetes/manifests` folder.

### Deploying through kubectl

You can deploy Jamadaar by running the following kubectl commands:

```bash
kubectl apply -f configmap.yaml -n <namespace>
kubectl apply -f rbac.yaml -n <namespace>
kubectl apply -f deployment.yaml -n <namespace>
```

### Helm Charts

Or alternatively if you configured `helm` on your cluster, you can deploy Jamadaar via helm chart located under `deployments/kubernetes/chart/Jamadaar` folder.

## Help

**Got a question?**
File a GitHub [issue](https://github.com/stakater/Jamadaar/issues), or send us an [email](mailto:stakater@gmail.com).

### Talk to us on Slack
Join and talk to us on the #tools-imc channel for discussing Jamadaar

[![Join Slack](https://stakater.github.io/README/stakater-join-slack-btn.png)](https://stakater-slack.herokuapp.com/)
[![Chat](https://stakater.github.io/README/stakater-chat-btn.png)](https://stakater.slack.com/messages/CA66MMYSE/)

## Contributing

### Bug Reports & Feature Requests

Please use the [issue tracker](https://github.com/stakater/Jamadaar/issues) to report any bugs or file feature requests.

### Developing

PRs are welcome. In general, we follow the "fork-and-pull" Git workflow.

 1. **Fork** the repo on GitHub
 2. **Clone** the project to your own machine
 3. **Commit** changes to your own branch
 4. **Push** your work back up to your fork
 5. Submit a **Pull request** so that we can review your changes

NOTE: Be sure to merge the latest from "upstream" before making a pull request!

## Changelog

View our closed [Pull Requests](https://github.com/stakater/Jamadaar/pulls?q=is%3Apr+is%3Aclosed).

## License

Apache2 © [Stakater](http://stakater.com)

## About

`Jamadaar` is maintained by [Stakater][website]. Like it? Please let us know at <hello@stakater.com>

See [our other projects][community]
or contact us in case of professional services and queries on <hello@stakater.com>

  [website]: http://stakater.com/
  [community]: https://github.com/stakater/
