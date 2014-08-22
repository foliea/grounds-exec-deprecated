## Grounds

This the official repository of the `Grounds` project.

This project is written in `Ruby` and `Go`.

To hack on this project, you need first to install `docker`.

Checkout the official `docker` [documentation](https://docs.docker.com/installation/)
to install it on your platform.

This project is made with three major parts:

- A web application.
- A websocket server application.
- A bunch of `docker` images to run code inside a container.

## Web application

The web application is written in `Ruby` with the web framework `Ruby on Rails`.

## Websocket application

The web application is written in `Go`.

## Docker images

There is one `docker` image for each language supported.

Checkout images [documentation](https://github.com/folieadrien/grounds/blob/master/docs/IMAGES.md)
to have more informations about how they are built and how to create your own.

## Contributing

To hack on `Grounds`, all you need is `git`, `make`, `docker` and your favorite text editor.

You can run the tests or the whole stack inside `docker` containers, with the same environment
used in production.

For instructions on setting up your development environment, please see this
[documentation](https://github.com/folieadrien/grounds/blob/master/docs/DEVENV.md).

Before sending a pull request, please checkout the contributing
[documentation](https://github.com/folieadrien/grounds/blob/master/docs/CONTRIBUTING.md).

## Licensing

`Grounds` is licensed under the MIT License. See `LICENSE` for full license text.
