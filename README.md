# go-api-starter
A good starting point for any Go API

## Getting Started
To run this api, simply run `make build && make run`, and it will launch the server with a health endpoint.

To use this as a starter, you can either use this as a template on GitHub or copy the code to your directory. First, remove the `go.mod` and `go.sum` files. Next, replace every instance of `bethanyj28/go-api-starter` with your repo name. Finally, run `go mod init && go mod tidy`.

## Inspiration
The folder layout was loosely inspired by Mat Ryer's [How I Write HTTP Services After Eight Years](https://pace.dev/blog/2018/05/09/how-I-write-http-services-after-eight-years.html). I really liked his thought about encapsulating handlers in functions that can perform setup and have a server struct to manage all the components of running a server. I took a lot of queues on layout from [golang-standards/project-layout](https://github.com/golang-standards/project-layout). It has a lot of great examples of real-world projects implementing those patterns, too<sup>*</sup>.

## Layout
I put the server struct and all handlers in `cmd/server` because I thought it made sense. It's nice compared to keeping it in the root because all app components are stored in their blocks. However, it does make the file structure more complicated. I won't be offended if you move it to the root if that's what you prefer. 

Any packages you write for your app that you would be okay with other applications using should go in `/pkg/{package_name}`. Any packages you write for your app that you want to keep private should go in `/internal/{package_name}`. Make sure to make interfaces for your packages to allow for easy testing. I highly recommend using [golang/mock](https://github.com/golang/mock) for mocking packages.

## Database
This template comes equipped with a postgres database using the [pgx](https://pkg.go.dev/github.com/jackc/pgx/v4) driver. It additionally uses [golang-migrate](https://pkg.go.dev/github.com/golang-migrate/migrate/v4) for migrations. However, it connects to the main application via an interface, so as long as a database store implements the interface you're golden. The app will not start up until migrations are added.

## TODO
- [x] Base API with health check
- [x] .env file setup
- [ ] Base tests
- [x] Database setup

<sup>*</sup>Since writing, apparently this is a controversial opinion whether the standards repo actually reflects good Go layout standards and is not endorsed by the Go team. As always, use your best judgement and never blindly copy code.
