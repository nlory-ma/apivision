package pigeon_phoenix
/*package main

import (
	"encoding/base64"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"net/http"
	"path/filepath"
	"github.com/julienschmidt/httprouter"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/vision/v1"
)


type ImageJson struct {

	Nom	[]string `json:"nom"`
	Date 	[]string `json:"date"`
	base64 []string `json:"base64"`

}




func init() {
    http.HandleFunc("/", HandleApi)
    http.HandleFunc("/api/appengine", ApiPost)
}

func static(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "client/"+r.URL.Path)
}

func HandleApi(w http.ResponseWriter, r *http.Request) {

	//instacie une nouvelle route
	Img = {

		Date:"12.1.2016"
		Nom: "chat.jpg"
		base64: "lfsoiheojkkreg9gh88yj9t85n5555f99e9t88555g5gg5fmpg56gkjhid9g5f5eke98"

	}


	router:= httprouter.New()
	switch req.Method {
		case "GET" :
			{
				router.GET("/api{image} ou  /https://pigeon-phoenix.appspot.com/", p httprouter.Params)
				var fichier = p.Byname("image")
				executeApi(fichier);
			}
		case "POST" :
		{
			var label = annotation.Vision;
			router.POST("/api/appengine", p httprouter.Params)
		}
}



//func init(){
//	http.HandleFunc("/", httpserver)
//	http.ListenAndServe(":4000", nil)
//}

//func httpserver(writer http.ResponseWriter, request *http.Request){

//	http.ServeFile(writer, request, "static/" + request.URL.Path)
//}:wq



// run submits a label request on a single image by given file.
// execute la requete pour une seule image dans un fichier


func executeApi(file string) error {
	ctx := context.Background()

	// Authenticate to generate a vision service.
	// authentification pour generer le service vision

	client, err := google.DefaultClient(ctx, vision.CloudPlatformScope)
	if err != nil {
		return err
	}
	service, err := vision.New(client)
	if err != nil {
		return err
	}

	// Read the image.
	// ouvrir le fichier et  lit l'image.
	// src := "gs://bucket_name/lenna.jpg" , definir la source en lui affectant l'url de l'image
	// Donc, c'est la src qu'on va ouvrir

	b, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Printf("notre erreur %d\n", 1)
		return err
	}

	// Construct a label request, encoding the image in base64.
	// construit une requete et encode l'image dans la base64.

	req := &vision.AnnotateImageRequest{
		Image: &vision.Image{
			Content: base64.StdEncoding.EncodeToString(b),
		},
		Features: []*vision.Feature{{Type: "LABEL_DETECTION"}},
	}
	batch := &vision.BatchAnnotateImagesRequest{
		Requests: []*vision.AnnotateImageRequest{req},
	}
	res, err := service.Images.Annotate(batch).Do()
	if err != nil {
		return err
	}

	// Parse annotations from responses
	// Parse les annotations de requetes
	if annotations := res.Responses[0].LabelAnnotations; len(annotations) > 0 {
		label := annotations[0].Description
		fmt.Printf("{ annotation :%s }\n", label)
		return nil
	}
	fmt.Printf("Not found label: %s\n", file)
	return nil
}


func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s <path-to-image>\n", filepath.Base(os.Args[0]))
	}
	flag.Parse()

	args := flag.Args()
	if len(args) == 0 {
		flag.Usage()
		os.Exit(1)
	}

	if err := executeApi(args[0]); err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		os.Exit(1)
	}
}*/






import (
    "net/http"
)

func init() {
    http.HandleFunc("/", static)
}

func static(w http.ResponseWriter, r *http.Request) {
    http.ServeFile(w, r, "client/"+r.URL.Path)
}


