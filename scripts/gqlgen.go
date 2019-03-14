// +build ignore

package main

import (
	"context"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/99designs/gqlgen/cmd"
	"github.com/google/go-github/github"
)

func main() {
	ctx := context.Background()
	dir := "../../../api"
	grdir := "../genresolver"
	grfile := filepath.Join(grdir, "/genresolver.go")
	// if generated resolver file already exists, remove it
	if _, err := os.Stat(grfile); err == nil {
		os.Remove(grfile)
	}
	// create temporary directory for genresolver if it doesn't exist
	if _, err := os.Stat(grdir); os.IsNotExist(err) {
		// 0777 denotes read, write, & execute for owner, group and others
		os.Mkdir(grdir, 0777)
	}
	// create temporary directory for schema if it doesn't exist
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.Mkdir(dir, 0777)
	}
	// remove temp directory and its contents at end of script
	defer func() {
		err := os.RemoveAll(dir)
		if err != nil {
			log.Fatalf("error removing temporary schema directory %s", err)
		}
		fmt.Printf("successfully removed temporary schema directory %s\n", dir)
	}()
	// initialize github client
	client := github.NewClient(nil)
	// get all files from our schema repo
	_, files, _, err := client.Repositories.GetContents(ctx, "dictyBase", "graphql-schema", "/", nil)
	if err != nil {
		log.Fatalf("error in getting graphql schema", err)
	}
	// loop through these files
	for _, n := range files {
		// get each individual file from our repo
		file, _, _, err := client.Repositories.GetContents(ctx, "dictyBase", "graphql-schema", n.GetName(), nil)
		if err != nil {
			log.Fatalf("error in getting individual schema file", err)
		}
		// decode file contents
		d, err := base64.StdEncoding.DecodeString(*file.Content)
		if err != nil {
			log.Fatalf("error decoding github file contents", err)
		}
		// create desired file path
		fp := filepath.Join(dir, "/", file.GetName())
		// write the github file content into its own local file
		err = ioutil.WriteFile(fp, d, 0777)
		if err != nil {
			log.Fatalf("error writing file", err)
		}
		fmt.Printf("successfully wrote file %s\n", fp)
	}
	cmd.Execute()
}
