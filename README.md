## Getting Started

Cloned copies of this repository must be placed in the correct location in the GOPATH

```
git clone git@github.com:chandrafortuna/assesment.git
```

### Running the application

Navigate to `assesment` folder and build and run it:

```
cd assesment/
go build -o assesment
./assesment
```

The following table shows the HTTP methods and URLs that represent the action supported in the API.

| Request  | Description |
| ------------- | ------------- |
| `POST /balls/init`  | Initial Setup container  |
| `POST /balls/addToContainer`  | Add ball into random container  |
| `GET /containers`  | Get all container  |
| `GET /containers/verified`  | Get verified container  |
| `GET /productVariant`  | Get All product variant  |
| `POST /productVariant`  | Add Product Variant  |
| `POST /productVariant/reserve¸`  | Reduce Qty from Product Variant  |
| `POST /homekey`  | Joni's Key Solution  |


For example request, please find and import assessment collection in the source code.
