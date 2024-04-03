# http_serve: A Simple HTTP Utility for Hosting Websites
http_serve is a small utility that allows you to quickly host websites for testing purposes directly from the shell. Whether you’re testing a web page, serving static files, or sharing content locally, http_serve simplifies the process.

## Installation

To install http_serve, follow these steps:

```shell
$ git clone https://github.com/SkillfulElectro/http_serve.git
$ make
```
## Usage

# Serve the Current Directory

To serve the contents of the current directory, simply run:

```shell
$ http_serve
```
# Serve Another Directory or File

If you want to serve files from a specific directory or a single file, provide the path as an argument:

```shell
$ http_serve File_or_dir_path
```

Replace File_or_dir_path with the actual path to the directory or file you want to serve.

### Extra Possibilities with http_serve

## GUI Applications and File Transfer:


    The concept behind http_serve can be extended to create GUI applications. By designing a user interface (UI) using web technologies (HTML, CSS, JavaScript), you can serve the UI directly from your shell.

    Additionally, this small utility can be handy for transferring files without the need for cables. Imagine the convenience of sharing files directly from your shell over a local area network (LAN)!


# File Transfer via LAN:

Suppose you have two devices connected via LAN. The sender can use http_serve to serve a specific file or an entire directory. The receiver can then connect to the website hosted by the sender and download any files they need from the LAN network.

# GUI Applications with Web Frontend and Backend:

    When creating GUI applications, you can use web technologies for the frontend (UI) and choose any backend technology you prefer.
    Here’s how it works:
        Design your UI using HTML, CSS, and JavaScript.
        Call http_serve to serve your UI on a local web server.
        Use websockets, REST APIs, or other communication methods to connect your frontend to the backend logic.
        The backend can be written in any language or framework of your choice (e.g., Python, Go, Node.js, etc.).


Enjoy using http_serve for quick and easy local web hosting! 🚀
