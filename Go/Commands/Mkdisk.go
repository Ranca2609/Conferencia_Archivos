package Commands

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func Create_Disk(unit string, path string, fit string, size string) {
	size_, err := strconv.Atoi(size)
	if err != nil {
		log.Fatal(err)
	}
	if path == "" {
		fmt.Println(red + "Error: Falta el parametro obligatorio Path, intentelo denuevo." + closing)
	} else if size_ <= 0 {
		fmt.Println(red + "Error: El Size no puede ser 0 o menor a 0, cambielo e intentelo denuevo." + closing)
	} else if FileExist(path) == true {
		fmt.Println(red + "Error: El disco " + yellow + find_last_of(path) + closing + red + " ya existe." + closing)
	} else {
		dt := time.Now()
		size_identified := new_size(unit, size_)
		ceros_ := make([]byte, size_identified)
		for i := 0; i < size_identified; i++ {
			ceros_[i] = 0
		}
		file, err := os.Create(path)
		if err != nil {
			fmt.Println(red + "Ocurrio un error inesperado, intente denuevo." + closing)
		}
		file.Write(ceros_)
		signature := DiskSignature()
		var bufferControl bytes.Buffer
		mbr_object := MBR{}
		copy(mbr_object.mbr_path[:], path)
		copy(mbr_object.mbr_date[:], dt.String())
		copy(mbr_object.mbr_size[:], strconv.Itoa(size_identified))
		copy(mbr_object.mbr_disk_signature[:], strconv.Itoa(signature))
		copy(mbr_object.mbr_fit[:], fit)
		for i := 0; i < 4; i++ {
			copy(mbr_object.mbr_particion[i].part_status[:], "0")
			copy(mbr_object.mbr_particion[i].part_type[:], "P")
			copy(mbr_object.mbr_particion[i].part_fit[:], fit)
			copy(mbr_object.mbr_particion[i].part_start[:], "0")
			copy(mbr_object.mbr_particion[i].part_size[:], "0")
			copy(mbr_object.mbr_particion[i].part_name[:], "")
		}
		file.Seek(0, os.SEEK_SET)
		mbr_in_bytes := Convert_to_Bytes(mbr_object)
		binary.Write(&bufferControl, binary.BigEndian, &mbr_in_bytes)
		WriteBytes(file, bufferControl.Bytes())
		file.Close()
		Print_Msg_Disk_Created(dt, path, size_identified, signature)
	}
}
