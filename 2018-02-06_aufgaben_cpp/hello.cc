// compile and execute:
// g++ hello.cc -o hello
// ./hello

#include <iostream>
using namespace std;

class Person {
  private:
    std::string name;
    std::string first;
    std::string last;
  public:
    Person(std::string first, std::string last);
    void set_name(std::string);
    bool name_first(std::string);
};

Person::Person(std::string first, std::string last) {
  this->first = first;
  this->last = last;
}

void Person::set_name (std::string name) {
  this->name = name;
}

bool Person::name_first(std::string first) {
  std::size_t found = name.find(first);
  if (found == std::string::npos) {
    return false;
  }
  // cout << "position:" << found << "\n";
  return true;
}

int main() 
{
  //
  // Zeichenkette initalizieren
  //

  std::string a = std::string("(a) hello");
  std::string * b = new std::string("(b) world");
  std::string gute_frage = std::string("Das ist eine gute Frage");

  cout << a;
  cout << "\n";

  cout << *b;
  cout << "\n";

  cout << gute_frage;
  cout << "\n\n";

  //
  // Zeichenkette suchen
  //

  std::string first = std::string("Aloys");
  Person * m_std = new Person("my first name", "my last name"); // constructor
  m_std->set_name("Kampinsky, Aloys ius");

  bool found = m_std->name_first(first);
  if (found) {
    cout << "Der Vorname '" << first << "' wurde gefunden\n";
  } else {
    cout << "Der Vorname '" << first << "' konnte NICHT gefunden werden\n";
  }
  cout << "\n";

  //
  // Gleitkommazahl berechnen
  //

  float   f1 = 1.0;
  double f2 = 1.0;
  long double f3 = 1.0;

  cout << "float:" << sizeof(f1) << " bytes\n";
  cout << "double:" << sizeof(f2) << " bytes\n";
  cout << "long double:" << sizeof(f3) << " bytes\n";
  cout << "(one byte is 8 bits)\n";

  return 0;
}


