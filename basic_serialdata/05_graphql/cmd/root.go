package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/neelance/graphql-go/relay"

	rs "github.com/muhfaris/goFun/basic_serialdata/05_graphql/pkg/resolvers"
	graphql "github.com/neelance/graphql-go"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	schema *graphql.Schema
)

var RootCmd = &cobra.Command{
	Use:   viper.GetString("app.name"),
	Short: "Exercise",
	Long:  "Learn use cobra library",
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("This 		: PreRun")
		fmt.Println("App name is 	:", viper.GetString("app.name"))
		fmt.Println("Version 	:", viper.GetString("app.version"))
	},
	Run: func(cmd *cobra.Command, args []string) {

		http.Handle("/", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			page, err := ioutil.ReadFile("graphiql.html")
			if err != nil {
				log.Fatal(err)
			}
			w.Write(page)
		}))
		http.Handle("/graphql", &relay.Handler{Schema: schema})
		log.Fatal(http.ListenAndServe(":8080", nil))
	},
}

func init() {
	cobra.OnInitialize(initSchema)
}

func initSchema() {
	schemaFile, err := ioutil.ReadFile("pkg/schema/schemahello.graphql")
	if err != nil {
		panic(err)
	}

	schema, err = graphql.ParseSchema(string(schemaFile), &rs.Resolvers{})
	if err != nil {
		panic(err)
	}
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
