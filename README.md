API wrapper for the Studio Ghibli API

An API that allows us to ​request a list of movies that feature a particular species​.

Example request: ​ http://localhost:8080​ /movies?species=<species_id>
Example response:
[
    {
        "id": "2baf70d1-42bb-4437-b551-e5fed5a87abe",
        "title": "Castle in the Sky",
        "description": "The orphan Sheeta ...",
        "director": "Hayao Miyazaki",
        "producer": "Isao Takahata",
        "release_date": "1986",
        "rt_score": "95"
    },
    {
        "id": "12cfb892-aac0-4c5b-94af-521852e46d6a",
        "title": "Grave of the Fireflies",
        "description": "In the latter part of World War II...",
        "director": "Isao Takahata",
        "producer": "Toru Hara","release_date": "1988",
        "rt_score": "97"
    }
]