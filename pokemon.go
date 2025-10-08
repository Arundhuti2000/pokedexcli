package main

// type pokemon struct {
// 	Name    string
// 	PokeUrl string
// }

// type pokemon_encounters struct {
// 	pokemon        pokemon
// 	versionDetails []byte
// }

type Pokemon_by_location_area struct {
	encounter_method_rate []byte             `json: "encounter_method_rate"`
	game_index            int                `json: "game_index"`
	id                    int                `json: "id"`
	location              string             `json: "location"`
	name                  string             `json: "name"`
	names                 string             `json: "names"`
	pokemon_encounters    []struct{
		Pokemon []struct{
			Name string `json:"name"`
			PokeUrl  string `json:"url"`
		}
		versionDetails []byte `json: "version_details"`
	} 
}