this contains all internal packages , and components

containing

- [] customers
- [] employees
- [] orders
- [] products
- [] routes
  - `constains all routes`
### routes needed
- employees (CRUD)
    -[x] create
    -[x] find
    -[] update role(s)
    -[x] delete
    -[] history (only for barista)
- products (CRUD)
    -[] create
    -[] find
    -[] update ( to update as a object )
    -[] delete        
- customers
    -[] create
    -[] find ( find order's by customer )
    -[] history ( order history )
- orders
    -[] create ( create order )
    -[] update ( update order )
    -[] find ( find order by customer )



## routes

- [] take/give order
    ```json
    PUT order/

    {
        "order":ORDER,
        "user":USER
    }
    ```
    - takes order from customer
        - order data
        - user data
- [] check order
    ```
    GET order/:id
    ```
    - check order status
- [] fullfill order
    ```json
    PATCH order/fullfill

    {  
        "product":PRODUCT.id
    }
    isBaked:status chenged then 
    ```
    - barista fullfills a part of order
    - if all product(s) of order are fullfilled
    - order status is changed to `baked`
- [] complete order
    ```json
    PATCH order/complete

    {
        "order":ORDER.id
    }
    ```
    - `baked` order taken by customer changed to `completed`
- [] add inventory
    ```json
    PUT product

    {
        qty:10,
        name:asjd,
        price:53
    }

    ```
    - manager adds `product` to inventory
- [] refreash inventory
    ```json
    PATCH product

    {   
        product_id:id,
        qty:num
    }
    ```
    - manager updated `product` inventory
- [] add employee
    ```json
    PUT employee

    {
        role:["role"],
        name:asjd,
        email:adsaf@gmail.com
        password:lkasjdlajsd
    }
    ```
    - manager add's barista or manager
- [] customer history
    ```json
    GET customer/history/:id
    ```
    - to check customer order history
    - add pagination
- [] barista  history
    ```json
    GET barista/history/:id
    ```
    - to check barista order fullfillment history
    - add pagination
