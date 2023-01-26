# Introduction

This is prototype backend for e-library that use `https://openlibrary.org/` API to fetch books data.<br> <br>
This app is built using `Go v1.18` and also have implemented unit tests using `testify` <br> <br>

Unit test results: <br>

1.  `/internal/core/service`
    ![Alt /core/service](./assets/service.png)
    <br><br>
2.  `/internal/repository`
    ![Alt /core/service](./assets/repository.png)
    <br><br>
3.  `/internal/utils`
    ![Alt /core/service](./assets/utils.png)
    <br><br><br>

# How to run

### Prerequisites:

    - Make sure you have `Golang` environment (Go v1.18)
    - Make sure you have `Make` environment
    - Make sure you have installed `postman`

1. Run command `make run-api` on your terminal at base directory of this project.
2. See `./joybox.postman_collection.json` to see how these API works.
