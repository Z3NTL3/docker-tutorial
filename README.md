### Services
- ``backend``
- ``caddy``

### Concept
- ``website``

    > Responsable to setup the frontend and backend server. It copies the required assets from the build context (project root directories) into the container. And additionally uses Ubuntu and installs nano (text editor). 
    >
    > It should be ensured that this service starts before ``caddy``.
    > Because the ``caddy`` services requires the frontend assets, a volume is created, which is shared between those. It will use this volume to gain frontend assets into it's own workspace.

- ``caddy``

    > Responsable to setup Caddy and uses the shared volume between ``website`` to setup frontend assets. It also has a bind mount with the host for the Caddyfile.
    > Backend is not known to Caddy but through a shared network between the containers they can interact with each other.

    >


**This more complex concept** ensures proper understanding of Docker Engine, Docker Compose, Containers, Volumes (bind mounts inclusive) and Networks in contrast to the basic tutorials.

<hr>

### Required: Setup ``user-defined`` bridge: _website2-caddy-bridge_
```docker network create website2-caddy-bridge```