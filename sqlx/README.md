## Macaron + sqlx

This is an example of using Macaron and sqlx to achieve CRUD tasks for songs.

- Create a song:

	```sh
	curl --data "title=Shake It Off&artist=Taylor Swift&year=2014" http://localhost:4000/songs
	```

- Get all songs:

	```sh
	curl http://localhost:4000/songs
	```

- Get a song:

	```sh
	curl http://localhost:4000/songs/1
	```

- Edit a song:

	```sh
	curl --request PATCH --data "title=Shake It Off2&artist=Taylor Swift&year=2014" http://localhost:4000/songs/1
	```

- Delete a song:

	```sh
	curl --request DELETE http://localhost:4000/songs/1
	```