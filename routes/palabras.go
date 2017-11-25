package routes

import (
	elastic "gopkg.in/olivere/elastic.v5"

	"github.com/gin-gonic/gin"

	"encoding/json"
	"net/http"

	"github.com/iorellana18/Team-Udes-Back/db"
	"github.com/iorellana18/Team-Udes-Back/utils"
)

type Palabras []Palabra

type Palabra struct {
	Palabra    string `json:"key"`
	Frecuencia int    `json:"doc_count"`
}

func PalabrasPorCategoria(c *gin.Context) {
	categoria := c.Query("id")

	ctx, client := db.ElasticInit()

	categoriaQuery := elastic.NewTermQuery("Categorias", categoria)
	termsAggregation := elastic.NewTermsAggregation().Field("Texto").ExcludeValues("a", "acá", "ahí", "ajena", "ajenas", "ajeno", "ajenos",
		"al", "algo", "algún", "alguna", "algunas", "alguno", "algunos", "allá", "alli", "allí", "ambos", "ampleamos", "ante", "antes",
		"aquel", "aquella", "aquellas", "aquello", "aquellos", "aqui", "aquí", "arriba", "asi", "atras", "aun", "aunque", "bajo", "bastante",
		"bien", "cabe", "cada", "casi", "cierta", "ciertas", "cierto", "ciertos", "como", "cómo", "con", "conmigo", "conseguimos", "conseguir",
		"consigo", "consigue", "consiguen", "consigues", "contigo", "contra", "cual", "cuales", "cualquier", "cualquiera", "cualquieras",
		"cuan", "cuán", "cuando", "cuanta", "cuánta", "cuantas", "cuántas", "cuanto", "cuánto", "cuantos", "cuántos", "de", "dejar", "del",
		"demás", "demas", "demasiada", "demasiadas", "demasiado", "demasiados", "dentro", "desde", "donde", "dos", "el", "él", "ella", "ellas",
		"ello", "ellos", "empleais", "emplean", "emplear", "empleas", "empleo", "en", "encima", "entonces", "entre", "era", "eramos", "eran",
		"eras", "eres", "es", "esa", "esas", "ese", "eso", "esos", "esta", "estaba", "estado", "estais", "estamos", "estan", "estar", "estas",
		"este", "esto", "estos", "estoy", "etc", "fin", "fue", "fueron", "fui", "fuimos", "gueno", "ha", "hace", "haceis", "hacemos", "hacen",
		"hacer", "haces", "hacia", "hago", "hasta", "incluso", "intenta", "intentais", "intentamos", "intentan", "intentar", "intentas",
		"intento", "ir", "jamás", "junto", "juntos", "la", "largo", "las", "lo", "los", "mas", "más", "me", "menos", "mi", "mía", "mia", "mias",
		"mientras", "mio", "mío", "mios", "mis", "misma", "mismas", "mismo", "mismos", "modo", "mucha", "muchas", "muchísima", "muchísimas",
		"muchísimo", "muchísimos", "mucho", "muchos", "muy", "nada", "ni", "ningun", "ninguna", "ningunas", "ninguno", "ningunos", "no", "nos",
		"nosotras", "nosotros", "nuestra", "nuestras", "nuestro", "nuestros", "nunca", "os", "otra", "otras", "otro", "otros", "para", "parecer",
		"pero", "poca", "pocas", "poco", "pocos", "podeis", "podemos", "poder", "podria", "podriais", "podriamos", "podrian", "podrias", "por",
		"por qué", "porque", "primero", "primero desde", "puede", "pueden", "puedo", "pues", "que", "qué", "querer", "quien", "quién", "quienes",
		"quienesquiera", "quienquiera", "quiza", "quizas", "sabe", "sabeis", "sabemos", "saben", "saber", "sabes", "se", "segun", "ser", "si",
		"sí", "siempre", "siendo", "sin", "sín", "sino", "so", "sobre", "sois", "solamente", "solo", "somos", "soy", "sr", "sra", "sres", "sta",
		"su", "sus", "suya", "suyas", "suyo", "suyos", "tal", "tales", "también", "tambien", "tampoco", "tan", "tanta", "tantas", "tanto", "tantos",
		"te", "teneis", "tenemos", "tener", "tengo", "ti", "tiempo", "tiene", "tienen", "toda", "todas", "todo", "todos", "tomar", "trabaja",
		"trabajais", "trabajamos", "trabajan", "trabajar", "trabajas", "trabajo", "tras", "tú", "tu", "tus", "tuya", "tuyo", "tuyos", "ultimo",
		"un", "una", "unas", "uno", "unos", "usa", "usais", "usamos", "usan", "usar", "usas", "uso", "usted", "ustedes", "va", "vais", "valor",
		"vamos", "van", "varias", "varios", "vaya", "verdad", "verdadera", "vosotras", "vosotros", "voy", "vuestra", "vuestras", "vuestro", "vuestros",
		"y", "ya", "yo", "http", "https", "www", "http", "t.co", "rt", "q")

	query, errSearch := client.Search(db.GetIndex()).Query(categoriaQuery).Aggregation("Palabras", termsAggregation).Do(ctx)
	utils.Check(errSearch)

	bucket, _ := query.Aggregations.BucketScript("Palabras")

	var palabras Palabras
	if bytesJson, err := bucket.Aggregations["buckets"].MarshalJSON(); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
	} else {
		if err := json.Unmarshal(bytesJson, &palabras); err != nil {
			c.String(http.StatusInternalServerError, err.Error())
		} else {
			c.JSON(http.StatusOK, palabras)
		}
	}
}
