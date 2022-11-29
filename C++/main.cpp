#include <iostream>
#include<bits/stdc++.h>
#include <vector>
#include "Commands/mkdisk.cpp"
#include "Commands/rmdisk.cpp"
using namespace std;


vector<string> split(string string_, char delimiter){
    vector<string> tmp_vector;
    string line;                                             
    stringstream input_stringstream(string_);                    
    while (getline(input_stringstream, line, delimiter))
    {
        tmp_vector.push_back(line);
    }
    return tmp_vector;
}

void Identify_Command(string line_){
    string size_,fit = "", unit = "", path = "", tipo = "", name = "";
    vector<string> incoming_line = split(line_, ' ');
    string command_ = incoming_line[0];
    transform(command_.begin(), command_.end(), command_.begin(), ::tolower);
    if (command_ == "mkdisk") {
        for(string tmp : incoming_line){
            vector<string> tmp_parameter = split(tmp, '=');
            transform(tmp_parameter[0].begin(), tmp_parameter[0].end(), tmp_parameter[0].begin(), ::tolower);
            if(tmp_parameter[0] == "-size"){size_ = tmp_parameter[1];}
            if(tmp_parameter[0] == "-unit"){unit = tmp_parameter[1];}
            if(tmp_parameter[0] == "-fit"){fit = tmp_parameter[1];}
            if(tmp_parameter[0] == "-path"){path = tmp_parameter[1];}
        }
        CreateDisc(path, size_, unit, fit);
    }else if (command_ == "rmdisk"){
        vector<string> tmp_parameter = split(incoming_line[1], '=');
        path = tmp_parameter[1];
        DeleteDisc(path);
    }else{
        std::cout << "Error: Comando no reconocido." << endl;
    }
}

int main () {
    Identify_Command("Comando entrante");
    return 0;
}

