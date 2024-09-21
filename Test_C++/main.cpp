#include <openssl/conf.h>
#include <openssl/evp.h>
#include <openssl/err.h>
#include <string.h>
#include <iostream>
#include <fstream>
#include <string>

using namespace std;

struct PasswordEntry {
    string website;
    string username;
    string password;
};

string encrypt(string text, int key) {
    for (int i = 0; i < text.length(); i++){
        text[i] = text[i] - key
    }
}