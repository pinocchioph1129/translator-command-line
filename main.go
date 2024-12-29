package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	bcp "translator-command-line/src/bcp"

	"cloud.google.com/go/translate"
	"golang.org/x/text/language"
)

func usage() {
	fmt.Println("Usage:")
	fmt.Println("-from <language> : Enter the language that you want to translate from")
	fmt.Println("-to <language> : Enter the language that you want to translate to")
}



func main () {
	var fromlanguage string
	var tolanguage string
	var decoy_fromlanguage string
	var decoy_tolanguage string
	var sentence string
	var exists bool

	flag.StringVar(&fromlanguage, "from", "", "Enter the language you want to translate from.")
	flag.StringVar(&tolanguage, "to", "", "Enter the language you want to translate to.")
	flag.Usage = func() {
		usage()
	}
	flag.Parse()
	if fromlanguage != "" {
		fromlanguage, exists = bcp.GetBCPTag(fromlanguage)
		if !exists {
			log.Fatalf("Initial 'from' language '%s' is not recognized.", fromlanguage)
		}
	}
	
	if tolanguage != "" {
		tolanguage, exists = bcp.GetBCPTag(tolanguage)
		if !exists {
			log.Fatalf("Initial 'to' language '%s' is not recognized.", tolanguage)
		}
	}
	
	scanner := bufio.NewScanner(os.Stdin);
	ctx := context.Background()
	client, err := translate.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()
	LOOP:
		for {
			
			fmt.Println("\nMenu:")
			fmt.Println("1. Change the language to translate from")
			fmt.Println("2. Change the language to translate to")
			fmt.Println("3. Translate a sentence")
			fmt.Println("4. Quit")
			fmt.Print("Enter your choice: ")
			if !scanner.Scan() {
				break
			}
			text := scanner.Text()
			if text == "" {
				break
			}
			
			switch text {
			case "1":
				fmt.Println("What language?")
				if scanner.Scan() {
					decoy_fromlanguage = scanner.Text()
					fromlanguage, exists = bcp.GetBCPTag(decoy_fromlanguage)
					if exists == false {
						fmt.Printf("Language '%s' is not recognized. Please try again.\n", decoy_fromlanguage)
						continue
					}
					fmt.Printf("The language is set to translate from: %s\n", decoy_fromlanguage)
				}
			case "2": 
			fmt.Println("What language?")
				if scanner.Scan() {
					decoy_tolanguage = scanner.Text()
					tolanguage, exists = bcp.GetBCPTag(decoy_tolanguage)
					if exists == false {
						fmt.Printf("Language '%s' is not recognized. Please try again.\n", decoy_tolanguage)
						continue				
					}
					fmt.Printf("The language is set to translate from: %s\n", decoy_tolanguage)
				}
			case "3":
				fmt.Println("Enter the sentence to translate:")
			if scanner.Scan() {
				sentence = scanner.Text()
				fromLangName := bcp.GetLanguageName(fromlanguage)
				toLangName := bcp.GetLanguageName(tolanguage)
				fmt.Printf("Translating '%s' from %s (%s) to %s (%s)...\n", 
					sentence, fromLangName, fromlanguage, toLangName, tolanguage)

				target, err := language.Parse(tolanguage)
				if err != nil {
					log.Fatalf("Failed to parse target language: %v", err)
				}
				resp, err := client.Translate(ctx, []string{sentence}, target, nil)
				if err != nil {
					log.Fatalf("Translation failed: %v", err)
				}
				fmt.Printf("Translated text: %v\n", resp[0].Text)
			}
			case "4":
				fmt.Println("Goodbye!")
				break LOOP
			default:
				fmt.Println("Invalid choice. Please select 1, 2, 3, or 4.")
			}
		}
}