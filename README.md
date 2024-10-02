# Idea
There should be a ``backend`` and ``frontend`` and a ``caddy`` instance.

``Caddy`` instance should manage serving a file server for the ``frontend``. But it should not be permitted to talk over the network towards the ``backend`` but only ``frontend``. 
However ``frontend`` can talk to both.

A volume between ``frontend`` and ``caddy`` is shared, don't ask why. Cuz why not. First, ``frontend`` gets the assets from the build context, which after ``caddy`` service copies the mount from ``frontend`` service.

Thats all.