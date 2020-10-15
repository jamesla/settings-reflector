# Settings-reflector

A container that reflects back settings, useful for integration tests.

Currently it just prints out environmental variables but in the future will add volume mounts etc.

Getting started:

```
docker run -p 8080:8080 -e HELLO=world jamesla/settings-reflector:$VERSION
curl http://localhost/env
```
