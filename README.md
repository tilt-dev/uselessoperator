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
    - Watch its status with `kubectl get machine -w` (if you make your terminal
      one line tall it almost looks like an animation)
5. The web controller makes the playful machine status accessible at
   localhost:8090
