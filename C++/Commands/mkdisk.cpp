#include <iostream>
#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <string.h>
#include <cstdlib>
#include <time.h>
#include "estructures.h"
using namespace std;


struct Particion
{
    char part_status;
    char part_type = 'p'; 
    char part_fit = 'f'; 
    int part_start;
    int part_size;
    char part_name[16];
};
struct MBR
{
    int mbr_tamano;
    time_t mbr_fecha_creacion;
    int mbr_disk_signature;
    string disk_fit;
    Particion mbr_partition[4];
};

int new_size(int size_, string unit){
    if (unit == "b"){
        return size_;
    }else if (unit == "k"){
        return size_ * 1024;
    }else if (unit == "m"){
        return size_ * 1048576;
    }
    return 0;
}



void CreateDisc(string path, string size_, string unit, string fit){
    string s = path;
    char sc[s.size() + 1];
    strcpy(sc, s.c_str());
    FILE *file=NULL;
    file=fopen(sc,"r");
    if(file!=NULL){
        cout<<"Ya existe el disco"<<endl;
        return;
    }
    int size_converted = stoi(size_);
    int tam = (new_size(size_converted, unit));;

    file=fopen(sc,"wb");
    fwrite("\0",1,1,file);
    fseek(file,tam,SEEK_SET); 
    fwrite("\0",1,1,file);
    MBR mbr;
    mbr.mbr_tamano = tam;
    mbr.mbr_disk_signature = rand()%1000;
    mbr.mbr_fecha_creacion = time(0);
    mbr.disk_fit = fit;
    for(int i = 0; i < 4; i++){
        mbr.mbr_partition[i].part_status = '0';
        mbr.mbr_partition[i].part_size = 0;
        mbr.mbr_partition[i].part_fit = 'f';
        mbr.mbr_partition[i].part_start = tam;
        strcpy(mbr.mbr_partition[i].part_name,"");
    }
    cout<<"-------------DISCO CREADO--------------------"<<endl;
    cout<<"Fecha de creacion: "<<asctime(gmtime(&mbr.mbr_fecha_creacion))<<endl;
    cout<<"Signature: "<<mbr.mbr_disk_signature <<endl;
    cout<<"Tamanio: "<<mbr.mbr_tamano <<endl;
    cout<<"Fit: " <<mbr.disk_fit <<endl;
    fseek(file,0,SEEK_SET);
    fwrite(&mbr,sizeof(MBR),1,file);
    fclose(file);
}