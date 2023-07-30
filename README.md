# Quakelog Challenge!

The goal of this project is to obtain data from the game Quake by parsing a log file that is located at the [root](https://github.com/viniciusfca/quakelog/blob/main/qgames.log) of the project.

# Technologies

## Golang

To use Golang, follow the installation instructions provided in the following [link](https://go.dev/doc/install).

# How to Run the Project

## Environment Variables

- [x] MONGO_URI
  - [x] **mongodb://mongo:27017**

## Running the Project

There are two ways to run the project:

1. At the root of the project, execute the command **go run main.go**
2. Run it using Docker Compose
   - For this, you need to install Docker on your machine. Installation instructions can be found in this [link](https://docs.docker.com/engine/install/).
   - After installing Docker, go to the root folder and execute the following command **docker compose up**

In both cases, the application will start on port **3000**.

## Endpoints

**[GET] - localhost:3000/v1/quake**
Endpoint responsible for returning all game data from the match.
**_Response 200_**

```json
[
  {
    "id": 2,
    "total_kills": 11,
    "players": ["Isgalamido", "Mocinha"],
    "kills": [
      {
        "player": "Isgalamido",
        "total": -5
      }
    ],
    "kills_by_means": {
      "MOD_FALLING": 1,
      "MOD_ROCKET_SPLASH": 3,
      "MOD_TRIGGER_HURT": 7
    }
  }
]
```

**Response 404**

```json
{
  "message": "Report not found for games"
}
```

**[GET] - localhost:3000/v1/quake/{gameId}**
Endpoint responsible for returning data for a specific game.
**_Response 200_**

```json
{
  "id": 4,
  "total_kills": 105,
  "players": ["Isgalamido", "Dono da Bola", "Zeh", "Assasinu Credi"],
  "kills": [
    {
      "player": "Zeh",
      "total": 20
    },
    {
      "player": "Isgalamido",
      "total": 19
    },
    {
      "player": "Dono da Bola",
      "total": 13
    },
    {
      "player": "Assasinu Credi",
      "total": 13
    }
  ],
  "kills_by_means": {
    "MOD_FALLING": 11,
    "MOD_MACHINEGUN": 4,
    "MOD_RAILGUN": 8,
    "MOD_ROCKET": 20,
    "MOD_ROCKET_SPLASH": 51,
    "MOD_SHOTGUN": 2,
    "MOD_TRIGGER_HURT": 9
  }
}
```

**Response 404**

```json
{
  "message": "Report not found for game ID : 4"
}
```

## Github Actions

The project includes two configured actions:

1. Executes unit tests for the application - [Action](https://github.com/viniciusfca/quakelog/blob/main/.github/workflows/go-test.yaml)
2. Builds the Docker image and updates the deployment file for Kubernetes - [Action](https://github.com/viniciusfca/quakelog/blob/main/.github/workflows/docker.yaml)

The necessary files for deploying the application in a Kubernetes environment can be found [here](https://github.com/viniciusfca/k8s/tree/main/quakelog).
Note: The image is only created and updated when a new release is created in the quakelog project.
