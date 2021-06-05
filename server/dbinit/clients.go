package dbinit

import (
	"log"
	"scootin/database"
	"scootin/global"

	"github.com/google/uuid"
)

/*----------------*/

func GenerateRandomClients() error {

	for _, name := range listOfRandomNames {

		row := database.RowType{
			"uuid": uuid.New(),
			"name": name,
		}

		_, err := global.DB.Insert("clients", row)
		if err != nil {
			log.Printf("\nError in data insertion: %v \nRow: \n%v", err, row)
		}

	}

	return nil
}

/*----------------*/

var listOfRandomNames []string = []string{
	"Kerrie Stockstill",
	"Stanley Caples",
	"Caryl Schmalz",
	"Angila Cadden",
	"Nora Schull",
	"Carol Summerall",
	"Daria Crumpton",
	"Chet Ryce",
	"Burl Havlik",
	"Shera Alfano",
	"Tien Rohrbaugh",
	"Zita Ervin",
	"Dayna Meredith",
	"Petra Velasques",
	"Jeanna Campoverde",
	"Joel Grabert",
	"Idella Boyster",
	"Coreen Tafolla",
	"Dora Eichler",
	"Micah Tullius",
	"Cyrstal Patenaude",
	"Adriana Spaulding",
	"Leonida Boughton",
	"Shonta Hagans",
	"Lyda Kates",
	"Reed Heaps",
	"Isela Gooslin",
	"Margie Dieterich",
	"Winford Dauber",
	"Mae Lytton",
	"Cole Escobedo",
	"Sade Siegrist",
	"Karri Pickron",
	"Roni Gourdine",
	"Donna Jeppson",
	"Lona Bayne",
	"Norbert Rotella",
	"Jacquie Britton",
	"Clementine Kozel",
	"Traci Fluellen",
	"Khalilah Wilhoit",
	"Jacquiline Robb",
	"Doria Barwick",
	"Galina Register",
	"Marion Steelman",
	"Brinda Dabrowski",
	"Susie Schwer",
	"Eloy Delisle",
	"Kristi Massenburg",
	"Ardis Delacruz",
	"Euna Condie",
	"Xenia Dangerfield",
	"Shenna Drown",
	"Kayleigh Pifer",
	"Tamiko Nolan",
	"Adrian Warthen",
	"Lyda Gaudin",
	"Delmar Buie",
	"Ulysses Mendel",
	"Ashlee Duquette",
	"Skye Horta",
	"Gary Heslin",
	"Emma Sitz",
	"Lisa Panetta",
	"Omar Rickard",
	"Ghislaine Arent",
	"Randall Thiessen",
	"Yadira Ferraro",
	"Solange Bartman",
	"Petronila Zapien",
	"Earnestine Obermiller",
	"Donnell Navin",
	"Michaele Cronin",
	"Myrtis Waltrip",
	"Merna Grieco",
	"Reatha Spahr",
	"Georgie Gryder",
	"Wally Stenzel",
	"Margeret Braverman",
	"Argentina Landgraf",
	"Sallie Naples",
	"Madge Buttrey",
	"Marx Agular",
	"Ollie Downard",
	"Nicholas Choudhury",
	"Robena Siggers",
	"Malinda Chuang",
	"Jenny Walkowiak",
	"Onie Gunnell",
	"Isobel Kircher",
	"Alix Dery",
	"Krystina Shawgo",
	"Ozell Vangieson",
	"Jamee Steinbeck",
	"Anissa Handler",
	"Ma Schepers",
	"Christiane Ringwood",
	"Adelaida River",
	"Kathie Lejeune",
	"Terese Oles",
} // Source: http://www.listofrandomnames.com
