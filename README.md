# flux2-bitbucketpipeline-dispatcher aka bpdispatcher

[![release](https://github.com/brainfair/flux2-bitbucketpipeline-dispatcher/actions/workflows/release.yml/badge.svg)](https://github.com/brainfair/flux2-bitbucketpipeline-dispatcher/actions/workflows/release.yml)

Bpdispatcher is a middleware server that receives a generic webhook from the fluxcd notification controller and triggers the bitbucket cloud pipeline with parameters from the received webhook

### Problem

Probably you want to promote HelmRelease version from staging to the production environment.
Flux provides [native support for GitHub dispatcher](https://fluxcd.io/flux/use-cases/gh-actions-helm-promotion/), but no options for bitbucket.

### Solver

With Bpdispatcher you can use Generic Webhook from Flux and trigger any bitbucket pipeline with the same automation.

* Install
  * [bpdispatcher install on Kubernetes](https://brainfair.github.io/flux2-bitbucketpipeline-dispatcher/)

* Bitbucket pipeline:
  * pipeline should be named pr-promotion under custom section ([pipeline template](bitbucket-pipeline.yml))

* FluxCD Provider and Alert:
  * [Example of Provide and Alert for flux](alert-example.yml)