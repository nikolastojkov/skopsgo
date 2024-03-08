# Skopsgo
An opinionated Go starter kit with Gin, HTMX, Tailwind, and Templ.

<img src="https://i.ibb.co/QHS4xb4/skopsgo-Github.png" height="480px" />

## Why
I needed a starter using HTMX, Tailwind, and Templ that I'll use for my own Go projects, and since this is how I usually structure my code I've decided to publish it to GitHub.

This is just what I use and is in no way perfect for every situation and/or project.

## Installation
You can use the starter either with Docker or standalone, though Docker is preferred.

Here I assume you have [Go](https://go.dev/) and [Docker](https://www.docker.com/) already installed.

You could/should rename the main module to your own project name with `go mod edit`

### Standalone Setup
1. Install the [Tailwind Standalone CLI](https://tailwindcss.com/blog/standalone-cli)
2. Compile style files `./tailwindcss -i ./web/static/main.css -o ./web/static/outputStyle.css`
3. Install Templ using `go install github.com/a-h/templ/cmd/templ@latest`
4. Generate the templates `templ generate`
5. Make a new env file `cp .env.example .env`
6. Build and/or run `cmd/skopsgo/skopsgo.go` or if you've renamed the modules, `cmd/MODULE_NAME/MODULE_NAME.go`

### Docker Setup
A Dockerfile is provided in order to build and run the app.

To build the container, run: `docker build --rm -t PROJECT_NAME .`

To run the container, run `docker run -p PORT:PORT PROJECT_NAME`

## How to use
After installing the starter you'll see that there is a splash screen with an example counter using HTMX.

This is to verify that everything is setup properly and working.

<img src="https://i.ibb.co/RhHDTRd/splash-Final.jpg" height="480px" />

You're safe to delete:
- The `splash.templ` file in `web/templates/splash.templ`
- The counter handlers at `internal/handlers/counter.go`
- Remove the `LoadCounterHandler` function call in `internal/handlers/handlers.go`

You're good to go with a blank canvas! Make a new template and serve it with your layout to start working on your pages.

## Future
Outside of maybe e2e testing and some other config files, I won't be updating this project much, since I'm using this as a bootstrap for all of my other projects, but you're free to request new things or submit PRs, and I'll see to them ASAP.
