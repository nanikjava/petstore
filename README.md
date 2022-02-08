This sample app shows one way to implement REST API for https://petstore.swagger.io/#/pet/findPetsByStatus.

The app shows the following:

* Using gorilla/mux to expose REST API
* Unit testing handler using `net/http/httptest`
* Processing are not using database but returning struct converted to JSON