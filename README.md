# e-comm-api-with-GO

<!-- TODO -->

### Building an E-commerce Api that will serve json response with some products.

<!-- Main Server -->

- [ ] set up Go project
- [ ] install Gorilla/mux: `go get -u github.com/gorilla/mux`
- [ ] design the json
  - [ ] each product will have: [{id, name, imageUrl, price, category}]
- [ ] create Products type struct
- [ ] connect with Firebase database

<!-- Front-end -->

- [ ] Create a display page where all products will be shown
  - [ ] fetch 10 products at a time to show on the page
  - [ ] each product will be shown in a product card
  - [ ] product card will have the image of the product and two buttons: one is for editing the product and the other one will have the delete product capability
  - [ ] Page will have a filter input field to filter by category
  - [ ] Page will have a pop up form to add one product
