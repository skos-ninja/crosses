# Noughts & Crosses

This is an example API built as part of a code test.
This game is designed to handle 2 users per game on a 3x3 grid.

This system is built in Go with a postgres backing.

The system has **NO** authentication on it and any requests take the `user_id` passed in as a true value

## Running

To run this system you will need an instance of postgres and set your connection string to an os environment variable named `PG_OPTS`.

Once you have a postgres database available you will need to initialise the tables by running each of the `.sql` scripts in the `db` folder.

Then simply run `go run cmd/crosses/crosses.go` and it will start and begin listening on port `8080`.

## Requests

### GET `/health`

Will return 204 if the system's database connection is healthy otherwise will return 500.

### POST `/create_player`

This will create a new user with a randomly generated ksuid. It's important to note that names are **not** unique.

#### Request

```json
{
    "name": "Jake"
}
```

#### Response

```json
{
    "id": "player_000000BegbfPRfhArzVXelKNNb1bU",
    "name": "Jake",
    "created_at": "2019-04-25T05:42:48.3049978+01:00"
}
```

### POST `/get_player`

Will return the details of that player based off of their id.

#### Request

```json
{
    "user_id": "player_000000BegbfPRfhArzVXelKNNb1bU"
}
```

#### Response

```json
{
    "id": "player_000000BegbfPRfhArzVXelKNNb1bU",
    "name": "Jake",
    "created_at": "2019-04-24T09:42:48.3049978Z"
}
```

### POST `/create_game`

This will create a new game and join the user passed in as player 1.

#### Request

```json
{
    "user_id": "player_000000BegbfPRfhArzVXelKNNb1bU"
}
```

#### Response

You should always expect the `player2_id` to be empty upon initial creation.

```json
{
    "id": "game_000000BegcnnmMBrwVeqMUrYGfgRM",
    "player1_id": "player_000000BegbfPRfhArzVXelKNNb1bU",
    "player2_id": "",
    "started_at": null,
    "finished_at": null,
    "winning_player": ""
}
```

### POST `/get_game`

Will return the game by ID.

#### Request

```json
{
    "game_id": "game_000000BegcnnmMBrwVeqMUrYGfgRM"
}
```

#### Response

```json
{
    "id": "game_000000BegcnnmMBrwVeqMUrYGfgRM",
    "player1_id": "player_000000BegbfPRfhArzVXelKNNb1bU",
    "player2_id": "",
    "started_at": null,
    "finished_at": null,
    "winning_player": ""
}
```

### POST `/list_available_games`

This will return any games that do not have all the players required to start and are available to join using `join_game`.

#### Response

```json
[
    {
        "id": "game_000000BegcnnmMBrwVeqMUrYGfgRM",
        "player1_id": "player_000000BegbfPRfhArzVXelKNNb1bU",
        "player2_id": "",
        "started_at": null,
        "finished_at": null,
        "winning_player": ""
   }
]
```

### POST `/join_game`

This will join the passed in user ID to the game and mark as ready to start.

#### Request

```json
{
	"game_id": "game_000000BegcgH9UVz2GWRMGvub8Y0e",
	"user_id": "player_000000BegbRSoTlcbIsGUv4Hx3Fuy"
}
```

#### Response

```json
{
    "id": "game_000000BegcgH9UVz2GWRMGvub8Y0e",
    "player1_id": "player_000000BegbfPRfhArzVXelKNNb1bU",
    "player2_id": "player_000000BegbRSoTlcbIsGUv4Hx3Fuy",
    "started_at": "2019-04-24T10:22:42.1910284Z",
    "finished_at": null,
    "winning_player": ""
}
```

### POST `/get_game_state`

This call will return all the game board positions with their current state.

#### Request

```json
{
    "game_id": "game_000000BegcgH9UVz2GWRMGvub8Y0e"
}
```

#### Response

`state` can be in three different states:

- `0` position not taken
- `1` player 1 position
- `2` player 2 position

```js
[
    {
        "x": 0,
        "y": 0,
        "state": 0
    }...
]
```

### POST `/take_turn`

This call will take a turn for the user at the given position.

#### Request

```json
{
    "game_id": "game_000000BegcgH9UVz2GWRMGvub8Y0e",
    "user_id": "player_000000BegbfPRfhArzVXelKNNb1bU",
    "x": 1,
    "y": 1
}
```

#### Response

If the turn taken causes a win then the game object will be updated with the winning player.

```json
{
    "game": {
        "id": "game_000000BegcgH9UVz2GWRMGvub8Y0e",
        "player1_id": "player_000000BegbfPRfhArzVXelKNNb1bU",
        "player2_id": "player_000000BegbRSoTlcbIsGUv4Hx3Fuy",
        "started_at": "2019-04-24T10:22:42.1910284Z",
        "finished_at": null,
        "winning_player": ""
    },
    "turn": {
        "id": "turn_000000Begvbhlp3CpGgRQmoCn0lHM",
        "game_id": "game_000000BegcgH9UVz2GWRMGvub8Y0e",
        "player_id": "player_000000BegbRSoTlcbIsGUv4Hx3Fuy",
        "x_axis": 0,
        "y_axis": 0,
        "created_at": "2019-04-24T15:48:58.7277554Z"
    }
}
```
