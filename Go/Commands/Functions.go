package Commands

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/cheggaaa/pb/v3"
)

var red string = "\033[1;31m"
var white string = "\033[1;37m"
var closing string = "\033[0m"
var yellow string = "\033[1;33m"
var blue string = "\033[1;34m"
var green string = "\033[1;32m"

func Create_Folder_IfNotExist(directorio string) {
	if _, err := os.Stat(directorio); os.IsNotExist(err) {
		err = os.Mkdir(directorio, 0755)
		if err != nil {
			panic(err)
		}
	}
}

func FileExist(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false
	} else {
		return true
	}
}

func Progress_Bar() {
	count := 200
	bar := pb.StartNew(count)
	for i := 0; i < count; i++ {
		bar.Increment()
		time.Sleep(time.Millisecond)
	}
	bar.Finish()
}

func WriteBytes(file *os.File, bytes []byte) {
	_, err := file.Write(bytes)
	if err != nil {
		log.Fatal(err)
	}
}

func Write_MRB(path string, mbr *MBR) {
	file_, err := os.OpenFile(string(path), os.O_RDWR, 0777)
	if err != nil {
		fmt.Println("Error: No se pudo abrir el archivo.")
	}
	newpos, err := file_.Seek(int64(0), os.SEEK_SET)
	if err != nil {
		fmt.Println("Error: No se pudo abrir el archivo.")
	}
	mbrByte := Convert_to_Bytes(mbr)

	_, _ = file_.WriteAt(mbrByte, newpos)
	file_.Close()
}

// Referencia de conversion de un Struct a Bytes y Viceversa https://gist.github.com/SteveBate/042960baa7a4795c3565
// Documentacion de Go sobre la libreria https://pkg.go.dev/encoding/gob

func Convert_Bytes_to_MBR(s []byte) MBR {
	p := MBR{}
	dec := gob.NewDecoder(bytes.NewReader(s))
	err := dec.Decode(&p)
	if err != nil && err != io.EOF {
		fmt.Println("Error: No se pudo hacer la conversion.")
	}
	return p
}

func Convert_to_Bytes(mbr interface{}) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(mbr)
	if err != nil && err == io.EOF {
		fmt.Println("Error: No se pudo hacer la conversion.")
	}
	return buf.Bytes()
}

func DiskSignature() int {
	Signature_Disk := rand.Intn(1000) * 1000
	return Signature_Disk
}

func find_last_of(path string) string {
	res1 := strings.Split(path, "/")
	nombre_disco := res1[len(res1)-1]
	return nombre_disco
}

func Print_Msg_Disk_Created(date_time time.Time, path string, size int, signature int) {
	fmt.Println(blue + "Disc created successfully" + closing)
	fmt.Println(green + "Name --->" + closing + yellow + find_last_of(path) + closing)
	fmt.Println(green + "Route --->" + closing + yellow + path + closing)
	fmt.Println(green + "Signature --->" + closing + yellow + strconv.Itoa(signature) + closing)
	fmt.Println(green + "Size --->" + closing + yellow + strconv.Itoa(size) + closing)
	fmt.Println(blue + "--------------------------------------------------" + closing)
}

func new_size(unit string, size int) int {
	unit_ := strings.ToLower(unit)
	if unit_ == "k" {
		new_size := size * 1024
		return new_size
	} else if unit_ == "b" {
		return size
	} else if unit_ == "m" {
		new_size := size * 1024 * 1024
		return new_size
	} else if unit_ == "" {
		new_size := size * 1024 * 1024
		return new_size
	} else {
		fmt.Println(red + "Error: Se ingreso un Unit no perteneciente al lenguaje." + closing)
	}
	return 0
}
