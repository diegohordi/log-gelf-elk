# log-gelf-elk
Elastic-Logstash-Kibana (ELK) - GELF logging 

This project contains the basic configuration of ELK for logging by the Docker GELF logging driver.

See: https://docs.docker.com/config/containers/logging/gelf/

# How to test
* Run `make run`
* Run `curl http://localhost:8081/stdout` and `curl http://localhost:8081/stderr`
* Open `http://localhost:5601/app/kibana`
* Set the index as `backend-*`



