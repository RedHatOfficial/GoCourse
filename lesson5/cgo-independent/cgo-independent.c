#include "cgo-independent.h"
const int eof = EOF;
int hello_from_c(char* hello){
  return printf("%s",hello);
}
