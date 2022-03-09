This is a golang REST api that reads and writes blog posts to a MongoDB database.

Here are the steps to run this after cloning on your local machine.
* docker build -t golang-rest-api .
* docker run -p 8080:8080 golang-rest-api

To create a blog post
* curl -d '{"title": "Hello", "text": "Hi there"}' http://localhost:8080/posts

To get all blog posts
* curl http://localhost:8080/posts