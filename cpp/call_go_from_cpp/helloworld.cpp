#include <iostream>
#include "libhelloplugin.h"

using namespace std;

int main() {
    GoString text = {"WOO!", 4};
    SayHello(text);

    return 0;
}
