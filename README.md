# flux2-bitbucketpipeline-dispatcher aka bpdispatcher

[![release](https://github.com/brainfair/flux2-bitbucketpipeline-dispatcher/actions/workflows/release.yml/badge.svg)](https://github.com/brainfair/flux2-bitbucketpipeline-dispatcher/actions/workflows/release.yml)

Bpdispatcher is a midleware server that receive generic webhook from fluxcd notification controller and trigger bitbucket cloud pipeline with parameters from received webhook

### Problem

Probaly you want to promote HelmRelease version from staging to production envorinment.
Flux provide [native support for github dispatcher](https://fluxcd.io/flux/use-cases/gh-actions-helm-promotion/), but no options for bitbucket.

### Solver

With Bpdispatcher you can use Generic Webhook from Flux and trigger any bitbucket pipeline with same automation.

* Install
  * [bpdispatcher install on Kubernetes](https://brainfair.github.io/flux2-bitbucketpipeline-dispatcher/)

* Bitbucket pipeline:
  * pipeline shoud be named pr-promotion under custom section ([pipeline template](bitbucket-pipeline.yml))

* FluxCD Provide and Alert:
  * [Example of Provide and Alert for flux](alert-example.yml)