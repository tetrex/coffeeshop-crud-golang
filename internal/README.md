this contains all internal packages , and components

containing

- [] customers
- [] employees
- [] orders
- [] products
- [] routes
  - `constains all routes`


## routes

- [] take/give order
    ```json
    POST order/

    {
        "order":ORDER,
        "user":USER
    }
    ```
    - takes order from customer
        - order data
        - user data
- [] check order
    - check order status
- [] fullfill order
    - barista fullfills a part of order
    - if all product(s) of order are fullfilled
    - order status is changed to `baked`
- [] complete order
    - `baked` order taken by customer changed to `completed`
- [] refreash inventory
    - manager updated `product` inventory
- [] add employee
    - manager add's barista
- [] customer history
    - to check customer order history
- [] barista  history
    - to check barista order fullfillment history
