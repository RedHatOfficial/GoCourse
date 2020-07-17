package main

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"io/ioutil"
)

type Item struct {
	Value string
}

func encodeStringIntoBSON(item Item) ([]byte, error) {
	bsonOutput, err := bson.Marshal(item)

	if err != nil {
		return bsonOutput, err
	} else {
		return bsonOutput, nil
	}
}

func encodeStringIntoJSON(item Item) ([]byte, error) {
	jsonOutput, err := json.Marshal(item)

	if err != nil {
		return jsonOutput, err
	} else {
		return jsonOutput, nil
	}
}

func encodeStringIntoIndentedJSON(item Item) ([]byte, error) {
	jsonOutput, err := json.MarshalIndent(item, "", "    ")

	if err != nil {
		return jsonOutput, err
	} else {
		return jsonOutput, nil
	}
}

func encodeStringIntoXML(item Item) ([]byte, error) {
	xmlOutput, err := xml.Marshal(item)

	if err != nil {
		return xmlOutput, err
	} else {
		return xmlOutput, nil
	}
}

func encodeStringIntoIndentedXML(item Item) ([]byte, error) {
	xmlOutput, err := xml.MarshalIndent(item, "", "    ")

	if err != nil {
		return xmlOutput, err
	} else {
		return xmlOutput, nil
	}
}

func encodeStringIntoGob(item Item) ([]byte, error) {
	var buffer bytes.Buffer
	encoder := gob.NewEncoder(&buffer)

	err := encoder.Encode(item)
	if err != nil {
		return buffer.Bytes(), err
	} else {
		return buffer.Bytes(), nil
	}
}

func saveString(encodedString []byte, filename string) {
	err := ioutil.WriteFile(filename, encodedString, 0644)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Stored into file", filename)
	}
}

func printBufferInfo(buffer []byte) {
	fmt.Println("\nBuffer with encoded string: ", len(buffer))
}

func main() {
	var item Item
	item.Value = `Další rozvoj různých forem činnosti vyzaduje rozšiřování logistických prostředků a nových návrhů. Pestré a bohaté zkušenosti jasně říkají, že konzultace se širokým aktivem dostatečně oddaluje propad odpovídajících podmínek aktivizace. Ideové úvahy nejvyššího řádu a rovněž stabilní a kvantitativní vzrůst a sféra naší aktivity ve značné míře podmiňuje vytvoření systému výchovy pracovníků odpovídajících aktuálním potřebám. Nesmíme však zapomínat, že navržená struktura organizace zvyšuje potřebu aplikace existujících finančních a administrativních podmínek. Poslání organizace, zejména pak stálé, informačně-propagandistické zabezpečení naší práce pomáhá udržovat kumulativní progresi pozic jednotlivých účastníků k zadaným úkolům. Tímto způsobem realizace plánovaných vytyčených úkolů vyvolává proces zavádění a modernizace systému masové účasti. Naše dlouhodobé ambice, stejně jako nový model organizační činnosti přetváří strukturu vedení směru progresivního rozvoje.`

	encodedString, err := encodeStringIntoXML(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedString)
	saveString(encodedString, "string1.xml")

	encodedString, err = encodeStringIntoIndentedXML(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedString)
	saveString(encodedString, "string2.xml")

	encodedString, err = encodeStringIntoJSON(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedString)
	saveString(encodedString, "string1.json")

	encodedString, err = encodeStringIntoIndentedJSON(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedString)
	saveString(encodedString, "string2.json")

	encodedString, err = encodeStringIntoBSON(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedString)
	saveString(encodedString, "string1.bson")

	encodedString, err = encodeStringIntoGob(item)
	if err != nil {
		fmt.Println(err)
		return
	}
	printBufferInfo(encodedString)
	saveString(encodedString, "string1.gob")
}
