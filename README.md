# Restaurant manager

A restaurant buyers' manager with DGraph Database, Golang API and Vue.js Interface
## Environment Variables

To run this project, you will need to add the following environment variables to your /api/.env file

`API_URL`

`DGRAPH_ENDPOINT` (if used as in `docker-compose.yml`, `db:9080` would work)

`API_PORT` (prefer `3000` to be aligned with the readme)
  
## Run Locally (Docker `18.06.0+` required!)

Clone the project

```bash
  git clone https://github.com/cristianblar/restaurant-manager
```

Go to the project directory

```bash
  cd restaurant-manager
```

Build docker-compose image

```bash
  docker-compose build
```

Start the whole project (it takes 90 seconds to ensure that links work properly through Docker containers)

```bash
  docker-compose up
```

Visit `http://localhost:5000` to interact directly with the inferface

Visit `http://localhost:3000` if you want to interact with the API from the browser, or consume the endpoints from any HTTP Requester...

## API Reference

#### Sync new date

```http
  GET /load-data
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `date` | `number` | **Required**. Date to sync into the DB, in Unix Timestamp format  |

#### Get all buyers (paginated by 100)

```http
  GET /api/buyers
```

| Parameter | Type     | Description                |
| :-------- | :------- | :------------------------- |
| `page` | `number` | The page you want to retrieve |

#### Get Buyer

```http
  GET /api/buyers/${id}
```

| Parameter | Type     | Description                       |
| :-------- | :------- | :-------------------------------- |
| `id`      | `string` | **Required**. Id of buyer to fetch |

## License

[MIT](https://choosealicense.com/licenses/mit/)
