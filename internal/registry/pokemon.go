package registry

// type pokemon struct {
// 	Name    string
// 	PokeUrl string
// }

// type pokemon_encounters struct {
// 	pokemon        pokemon
// 	versionDetails []byte
// }

type Pokemon_by_location_area struct {
	Encounter_method_rate []byte             `json:"encounter_method_rate"`
	Game_index            int                `json:"game_index"`
	Id                    int                `json:"id"`
	Location              string             `json:"location"`
	Name                  string             `json:"name"`
	Names                 string             `json:"names"`
	Pokemon_encounters    []struct{
		Pokemon struct{
			Name string `json:"name"`
			PokeUrl  string `json:"url"`
		}
		VersionDetails []byte `json:"version_details"`
	} 
}