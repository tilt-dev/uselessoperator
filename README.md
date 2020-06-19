# Useless Machine Operator

This is meant to replicate a useless machine (see YouTube) as a Kubernetes
operator. It's meant to show how to use KubeBuilder to:

- Create a CRD
- Create a controller
- Get and set fields in a resource
- Get, create, and delete resources
- Use a controller to watch a different resource type

In practice, here's what it does:

1. A machine resource can have three types: useful, useless, or playful
2. A useful machine is allowed to exist
3. A useless machine gets immediately deleted
4. A playful machine goes towards the off switch, comes back, goes
   again (etc.) until it finally turns itself off
    - Watch its status with `watch -n0.5 kubectl get machine`

To run it, install [Tilt](https://tilt.dev/) and `tilt up`.

Relevant links:

- [Best practices for building Kubernetes Operators and stateful
  apps](https://cloud.google.com/blog/products/containers-kubernetes/best-practices-for-building-kubernetes-operators-and-stateful-apps)
- [KubeBuilder](https://github.com/kubernetes-sigs/kubebuilder)
- Useless machine videos: [here](https://www.youtube.com/watch?v=aqAUmgE3WyM)
  and [here](https://www.youtube.com/watch?v=kproPsch7i0)
- [Tilt](https://tilt.dev/)
- [Printer Columns](https://book.kubebuilder.io/reference/generating-crd.html) for KubeBuilder
