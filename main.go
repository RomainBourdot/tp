package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Character struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Status  string `json:"status"`
	Species string `json:"species"`
	Type    string `json:"type"`
	Gender  string `json:"gender"`
	Image   string `json:"image"`
}

func FetchCharacters() ([]Character, error) {
	resp, err := http.Get("https://rickandmortyapi.com/api/character")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result struct {
		Results []Character `json:"results"`
	}
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}

	return result.Results, nil
}

func FetchCharacter(id int) (Character, error) {
	resp, err := http.Get(fmt.Sprintf("https://rickandmortyapi.com/api/character/%d", id))
	if err != nil {
		return Character{}, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Character{}, err
	}

	var character Character
	err = json.Unmarshal(body, &character)
	if err != nil {
		return Character{}, err
	}

	return character, nil
}

func main() {

	characters, err := FetchCharacters()
	if err != nil {
		fmt.Println("Erreur lors de la récupération des personnages:", err)
		return
	}
	fmt.Println("Liste des personnages:")
	for _, character := range characters {
		fmt.Printf("ID: %d, Nom: %s, Statut: %s, Espèce: %s\n", character.ID, character.Name, character.Status, character.Species)
	}

	character, err := FetchCharacter(1)
	if err != nil {
		fmt.Println("Erreur lors de la récupération du personnage:", err)
		return
	}
	fmt.Printf("\nDétails du personnage ID 1:\nNom: %s\nStatut: %s\nEspèce: %s\nGenre: %s\n", character.Name, character.Status, character.Species, character.Gender)
}
