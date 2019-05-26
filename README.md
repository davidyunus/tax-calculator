# Tax Calculator

This document provide tax calculator API for calculating tax

## List of Tax Calculator API

### Request and Response Examples

* #### Create tax data:

    POST - http://127.0.0.1:9090/v1/tax
    
    Request body :
    
    ```
    {
        "name": "holy cow",
        "taxCode": 1,
        "price": 5000
    }
    ```

* #### List all taxes:

    POST - http://127.0.0.1:9090/v1/taxes

    Response body :
    ```
    [
        {
            "taxId": 2,
            "name": "holy cow",
            "taxCode": 1,
            "type": "food and beverage",
            "refundable": "yes",
            "price": 5000,
            "tax": 500,
            "amount": 5500,
            "createdAt": "2019-05-26T18:16:30.376838Z",
            "updatedAt": "2019-05-26T18:16:30.376838Z",
            "deletedAt": null
        },
        {
            "taxId": 1,
            "name": "john wick",
            "taxCode": 3,
            "type": "entertainment",
            "refundable": "no",
            "price": 150,
            "tax": 0.5,
            "amount": 150.5,
            "createdAt": "2019-05-26T18:13:37.466326Z",
            "updatedAt": "2019-05-26T18:13:37.466326Z",
            "deletedAt": null
        },
        {
            "taxId": 3,
            "name": "gudang micin",
            "taxCode": 2,
            "type": "tobacco",
            "refundable": "no",
            "price": 1000,
            "tax": 30,
            "amount": 1030,
            "createdAt": "2019-05-26T18:16:52.450389Z",
            "updatedAt": "2019-05-26T18:16:52.450389Z",
            "deletedAt": null
        }
    ]
    ```

* #### Search tax data by `query`:

    GET - http://127.0.0.1:9090/v1/tax?query=holy

    Response body :
    ```
    [
        {
            "taxId": 2,
            "name": "holy cow",
            "taxCode": 1,
            "type": "food and beverage",
            "refundable": "yes",
            "price": 5000,
            "tax": 500,
            "amount": 5500,
            "createdAt": "2019-05-26T18:16:30.376838Z",
            "updatedAt": "2019-05-26T18:16:30.376838Z",
            "deletedAt": null
        }
    ]
    ```

* #### Search tax data by `taxId`:

    * GET - http://127.0.0.1:9090/v1/tax/2

    Response body :
    ```
    {
        "taxId": 2,
        "name": "holy cow",
        "taxCode": 1,
        "type": "food and beverage",
        "refundable": "yes",
        "price": 5000,
        "tax": 500,
        "amount": 5500,
        "createdAt": "2019-05-26T18:16:30.376838Z",
        "updatedAt": "2019-05-26T18:16:30.376838Z",
        "deletedAt": null
    }
    ```

* #### Update tax data :

    * PUT http://127.0.0.1:9090/v1/tax/{taxId}

    Request body :
    ```
    {
        "name": "burger king",
        "taxCode": 1,
        "price": 2000
    }
    ```
