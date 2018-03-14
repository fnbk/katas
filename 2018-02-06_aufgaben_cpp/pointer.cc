// compile and execute:
// g++ pointer.cc -o pointer
// ./pointer

#include <iostream>
using namespace std;

// Frage: Wozu benötigen wir einen NULL-Zeiger?
// 
// 1) Ein NULL-Zeiger zeigt auf eine ungültige Adresse.
//
// 2) Beispiel:
// double* ptr_x = NULL; // hier wir explizit angegeben, worauf der Pointer zeigt (nämlich auf einen ungültigen Bereich)
// double* ptr_x; // hier ist unsicher, auf welche Adresse der Pointer zeigt
//
// Wenn bei der Initialisierung keine explizite Adresse angegeben ist, zeigt der Pointer auf eine beliebige Adresse. Dies könnte später zu unerwarteten Speicherzugriffen führen.

int main() 
{
  double x; // initialisiert x mit undefiniertem Wert
  cout << "value of x:     " << x << std::endl;
  double* ptr_x = NULL; // initialisiert pointer mit expliziter Adresse (NULL)
  cout << "value of ptr_x: " << ptr_x << std::endl;
  ptr_x = &x; // referenziert Variable x (weist die Adresse der Variable x ptr_x zu)
  cout << "value of ptr_x: " << ptr_x << std::endl;
  x = *ptr_x; // dereferenziert Pointer (liest den an der Adresse liegenden Wert aus)
  if(!ptr_x) cout << "Warning: ptr_x does not have a meaningful address"; // prüft auf NULL
  return 0;
}

// output:
// value of x:     6.95322e-310   // jedes mal ein anderer Wert (undefiniert)
// value of ptr_x: 0x0            // zeigt immer auf NULL
// value of ptr_x: 0x7fff5770d850 // jedesmal eine andere Adresse (undefiniert)
