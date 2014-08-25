# </\> Grounds

[ ![Codeship Status for folieadrien/grounds](https://codeship.io/projects/19c7a0d0-f6e9-0131-f838-6aed31f6fcc8/status?branch=master)]
(https://codeship.io/projects/28538)

This the official repository of the `Grounds` project.

`Grounds` is a web application that compiles and executes arbitrary code and returns the program output.

`Grounds` support many languages and make it really trivial to add support for any language.

This project is written in `Ruby` and `Go`.

This project is made with three major parts:

- A web application.
- A websocket server application.
- A bunch of `docker` images to run code inside a container.

## Web application

The web application is written in `Ruby` with the awesome framework `Ruby on Rails`.

## Websocket application

The websocket server application is written in `Go`.

## Docker images

There is one `docker` image for each language supported.

Checkout images [documentation](https://github.com/folieadrien/grounds/blob/master/docs/IMAGES.md)
to have more informations about how they are built and how to create your own.

## Contributing

For instructions on setting up your development environment, please see this
[documentation](https://github.com/folieadrien/grounds/blob/master/docs/DEVENV.md).

Before sending a pull request, please checkout the contributing
[documentation](https://github.com/folieadrien/grounds/blob/master/docs/CONTRIBUTING.md).

## Licensing

`Grounds` is licensed under the MIT License. See `LICENSE` for full license text.
