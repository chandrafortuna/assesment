# Simple Messaging API

API for sending and retrieve message. 

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

Make a POST request to the /balls/init endpoint 
Request Body:
```
{
	"numContainer": 2, //specify number container used
	"maxLoad": 10 // maximum ball capacity of each container
}
```


## Running the tests

```
go test ./test/
```
