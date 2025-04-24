## clients

This package provides a client for interacting with a store API. It defines a StoreClient struct that allows users to set, get, delete, and find data in the store. The client uses the encoding/json package to marshal and unmarshal JSON data, the net/http package to make HTTP requests to the store API, the io package to read the response body from the store API, the bytes package to create a buffer for the JSON data, and the fmt package for error handling.

```
core/clients/store.go
```

The StoreClient struct has methods for setting, getting, deleting, and finding data in the store. The Set method takes a SetRequest struct as input and sends a POST request to the "stores/set" endpoint of the store API. The Get method takes a GetRequest struct as input and sends a POST request to the "stores/get" endpoint of the store API. It then parses the response and returns a GetResponse struct. The Delete method takes a DeleteRequest struct as input and sends a POST request to the "stores/delete" endpoint of the store API. The Find method takes a FindRequest struct as input and sends a POST request to the "stores/find" endpoint of the store API. It then parses the response and returns a FindResponse struct.

The code defines several structs to represent the data exchanged with the store API. These include SetRequest, GetRequest, GetResponse, DeleteRequest, FindRequest, and FindResponse. The SetRequest struct contains the keys and values to be set in the store. The GetRequest struct contains the keys to be retrieved from the store. The GetResponse struct contains the keys and values retrieved from the store. The DeleteRequest struct contains the keys to be deleted from the store. The FindRequest struct contains the topk and key to be used in the search. The FindResponse struct contains the keys, values, and similarities found in the search.

The doRequest and doRequestWithResponse helper functions are used to perform requests to the store API.

