# e-comm-api-with-GO

<!-- TODO -->

### Building an E-commerce Api that will serve json response with some products.

<!-- Main Server -->

- [ ] Set up Go project
- [ ] Install Gorilla/mux: `go get -u github.com/gorilla/mux`
- [ ] Design the json
  - [ ] Product will have: [{id, name, imageUrl, price, category}]
- [ ] Create products type struct
- [ ] Connect with Firebase database

<!-- Front-end -->

- [ ] Create a display page where all products will be shown
  - [ ] Fetch 10 products at a time to show on the page
  - [ ] Product will be shown in a product card
  - [ ] Product card will have the image of the product and two buttons: one is for editing the product and the other one will have the delete product capability
  - [ ] Page will have a filter input field to filter by category
  - [ ] Page will have a pop up form to add one product
