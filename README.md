## Usage

[Helm](https://helm.sh) must be installed to use the charts.  Please refer to
Helm's [documentation](https://helm.sh/docs) to get started.

Once Helm has been set up correctly, add the repo as follows:

    helm repo add bpdispatcher https://brainfair.github.io/flux2-bitbucketpipeline-dispatcher

If you had already added this repo earlier, run `helm repo update` to retrieve
the latest versions of the packages.  You can then run `helm search repo
bpdispatcher` to see the charts.

To install/upgrade the KubeLinks chart:

    helm upgrade bpdispatcher bpdispatcher/bpdispatcher -i

To uninstall the chart:

    helm delete bpdispatcher