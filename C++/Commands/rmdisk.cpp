#include <iostream>
#include <string.h>



void DeleteDisc(std::string path){
    char sc[path.size() + 1];
    strcpy(sc, path.c_str());
    remove(sc);
}