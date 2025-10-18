package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strconv"
)

type Article struct {
	Id          int
	Nom         string
	Prix        float64
	Image       string
	Reduction   float64
	Description string
	EstenReduc  bool
}

var articles = []Article{
	{Id: 1, Nom: "PALACE PULL A CAPUCHE UNISEXE CHASSEUR", Prix: 100, Image: "/static/img/products/19A.webp", Reduction: 75, Description: "Un pull tres sexy", EstenReduc: true},
	{Id: 2, Nom: "PALACE PULL A CAPUCHON MARINE", Prix: 120.99, Image: "/static/img/products/21A.webp", Reduction: 105.99, Description: "Un pull sexy+ et tiens chaud", EstenReduc: true},
	{Id: 3, Nom: "PALACE PULL CREW PASSEPOSE NOIR", Prix: 125.99, Image: "/static/img/products/22A.webp", Reduction: 105.99, Description: "Un pull sexy+ et tiens chaud", EstenReduc: true},
	{Id: 4, Nom: "PALACE WASCAPUCHON MARINEHED TERRY 1/4 PLACKET HOOD MOJITO", Prix: 150, Image: "/static/img/products/16A.webp", Reduction: 150, Description: "Un pull sexy++ et tiens chaud et jamais froid", EstenReduc: false},
	{Id: 5, Nom: "PALACE PANTALON BOSSY JEAN STONE", Prix: 80, Image: "/static/img/products/34B.webp", Reduction: 60, Description: "Un pantalon super sexy+ et tiens chaud pour l'hiver", EstenReduc: true},
	{Id: 6, Nom: "PALACE PANTALON CARGO GORE-TEX R-TEK NOIR", Prix: 100.99, Image: "/static/img/products/33B.webp", Reduction: 100.99, Description: "Un pantalon super sexy++ et tiens chaud++", EstenReduc: false},
}

func main() {
	listTemplates, errtemplates := template.ParseGlob("./templates/*.html")
	if errtemplates != nil {
		fmt.Println(errtemplates)
		os.Exit(1)
	}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		listTemplates.ExecuteTemplate(w, "index", articles)
	})

	http.HandleFunc("/produit", func(w http.ResponseWriter, r *http.Request) {
        idProduit := r.FormValue("Id")
        produitId, err := strconv.Atoi(idProduit)
        if err != nil {
            http.Error(w, "Erreur: id du produit invalide", http.StatusBadRequest)
            return
        }

        for _, product := range articles{
            if product.Id == produitId {
                listTemplates.ExecuteTemplate(w, "produit", product)
                return
            }
        }

        http.Error(w, "Produit non trouvé", http.StatusNotFound)
    })

	http.HandleFunc("/add-produit", func(w http.ResponseWriter, r *http.Request) {
		listTemplates.ExecuteTemplate(w, "add-produit", articles)
	})

	FileServer := http.FileServer(http.Dir("./assets"))
	http.Handle("/static/", http.StripPrefix("/static/", FileServer))

	fmt.Println("http://localhost:8000/")
	http.ListenAndServe("localhost:8000", nil)

}
