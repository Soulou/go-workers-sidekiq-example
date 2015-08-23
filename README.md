# Example of integration between go-workers and Sidekiq

# Sidekiq Worker

```
bundle
bundle exec sidekiq -r ./ruby_worker/hello_worker.rb
```

# Sidekiq Client

```
ruby ruby_client_schedule/client.rb
```

# Go Client/Server

```
go get github.com/Scalingo/go-workers
go run <path>
```
