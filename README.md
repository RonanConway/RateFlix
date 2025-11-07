# RateFlix

Writing microservice application in Golang for learning purposes as I'm new to Go.

Using HashiCorp Consul for Service Discovery

The command runs Hashicorp Consul inside Docker in development mode, exposing
its ports 8500 and 8600 for local use.

```
docker run -d -p 8500:8500 -p 8600:8600/udp --name=dev-consul
consul agent -server -ui -node=server-1 -bootstrap-expect=1
-client=0.0.0.0
```

Now, go to the Consul web UI via its link, http://localhost:8500/. When you open
the Services tab, you should see the list of services and an active Consul instances
