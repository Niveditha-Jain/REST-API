package main

import (
	"github.com/drone/routes"
	"log"
	"net/http"
	"encoding/json"
	"io/ioutil"
)


var profileMap map[string]profile = make(map[string]profile)

type profile struct {
	Email string `json:"email"`
	Zip string `json:"zip"`
	Country string `json:"country"`
	Profession string `json:"profession"`
	FavoriteColor string `json:"favorite_color"`
	IsSmoking string `json:"is_smoking"`
	FavoriteSport string `json:"favorite_sport"`
	Food  struct {
			  Type string `json:"type"`
			  DrinkAlcohol string `json:"drink_alcohol"`
		  } `json:"food"`
	Music struct {
			  SpotifyId string `json:"spotify_user_id"`
		  }`json:"music"`
	Movie struct{
           TVShows []string `json:"tv_shows"`
		   Movies []string `json:"movies"`
		  }`json:"movie"`
	Travel struct {
			  Flight struct{
						 Seat string `json:"seat"`
					 }`json:"flight"`
		  }`json:"travel"`
}





func main() {
	mux := routes.New()

	mux.Get("/profile/:email", GetProfile)
	mux.Post("/profile",PostProfile)
	mux.Put("/profile/:email", PutProfile)
    mux.Del("/profile/:email",DeleteProfile)
	http.Handle("/", mux)
	log.Println("Listening...")
	http.ListenAndServe(":3000", nil)


}



func GetProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("Get")
	params := r.URL.Query()
	email:= params.Get(":email")
	if getProfile, ok := profileMap[email]; ok {
		//Profile, _ := json.MarshalIndent(getProfile,"","    ")
		Profile, _ := json.Marshal(getProfile)
		// Convert bytes to string.
		log.Println(string(Profile))
		w.WriteHeader(200)
		w.Write([]byte(string(Profile)))
	}else{
		w.WriteHeader(404)
       w.Write([]byte("Email Not found"))
	}
}


func PostProfile(w http.ResponseWriter, r *http.Request){
	log.Println("Post")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}
	var p profile
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Fatal(err)
	}
 	profileMap[p.Email]= p
	Profile, _ := json.Marshal(p)
	// Convert bytes to string.
	log.Println(string(Profile))
	w.WriteHeader(201) // When the resource is successfully created, we use 201 status.
	w.Write([]byte(p.Email + "'s profile created"))
}

func DeleteProfile(w http.ResponseWriter, r *http.Request) {

	log.Println("Delete")

	params := r.URL.Query()
	email := params.Get(":email")
	if getProfile, ok := profileMap[email]; ok {
		delete(profileMap,getProfile.Email)
		w.WriteHeader(http.StatusNoContent)
	}else{
		w.Write([]byte("Email Not found"))
	}
}



func PutProfile(w http.ResponseWriter, r *http.Request) {
	log.Println("Put")
	params := r.URL.Query()
	email := params.Get(":email")
	if getProfile, ok := profileMap[email]; ok {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Fatal(err)
		}
		var q profile
		err = json.Unmarshal(body, &q)
		if err != nil {
			log.Fatal(err)
		}

		if(q.Email!=""){
			getProfile.Email=q.Email
		}
		if(q.Zip!=""){
			getProfile.Zip=q.Zip
		}
		if(q.Country!=""){
			getProfile.Country=q.Country
		}
		if(q.Profession!=""){
			getProfile.Profession=q.Profession
		}
		if(q.FavoriteColor!=""){
			getProfile.FavoriteColor=q.FavoriteColor
		}
		if(q.IsSmoking!=""){
			getProfile.IsSmoking=q.IsSmoking
		}
		if(q.FavoriteSport!=""){
			getProfile.FavoriteSport=q.FavoriteSport
		}
		if(q.Food.Type!=""){
			getProfile.Food.Type=q.Food.Type
		}
		if(q.Food.DrinkAlcohol!=""){
			getProfile.Food.DrinkAlcohol=q.Food.DrinkAlcohol
		}
		if(q.Music.SpotifyId!=""){
			getProfile.Music.SpotifyId=q.Music.SpotifyId
		}
		if(len(q.Movie.TVShows)!=0 ){
			getProfile.Movie.TVShows=q.Movie.TVShows
		}
		if(len(q.Movie.Movies)!=0){
			getProfile.Movie.Movies=q.Movie.Movies
		}
		if(q.Travel.Flight.Seat!=""){
			getProfile.Travel=q.Travel
		}
		profileMap[email]=getProfile;
		w.WriteHeader(204)
	}else {
		w.Write([]byte("Email Not found"))
	}

}







