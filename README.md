## Resource Autoscaler Operator

### A production-ready Kubernetes operator that monitors pod resource usage and automatically scales deployments or adjusts resource limits to prevent OOMKills and optimize resource utilization.


``` kubebuilder init --domain baniasadi.autoscaling.io --repo github.com/sanazba/resource-autoscaler-operator
``` kubebuilder create api --group autoscaling --version v1 --kind ResourceScalerPolicy --resource --controller

