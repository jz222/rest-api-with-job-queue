# REST API with RabbitMQ worker queue

A containerized worker setup with RabbitMQ as job queue. The API receives new messages and adds them to the queue. RabbitMQ manages the messages and sends them to the worker. Once the worker acknowledges that the work is completed, it will receive a new message from the queue. The queue is persisted so that messages aren't lost after a container restart. In this setup, there is only one worker container. However, additional worker containers can be added easily to handle larger workloads.

## Usage

Pull the repository and add the corresponding `.env` files to the two services. Install dependencies of each of the service with `go get`. Afterwards, start the containers with `docker-compose up --build`.

## Enqueue Jobs

Once the containers are running you can enqueue jobs by sending a `POST` request to the API at `http://localhost:2500/addjob` with the following body:

```
{
	"message": "Hello World!"
}
```

The RabbitMQ dashboard is available at `http://localhost:2800`.