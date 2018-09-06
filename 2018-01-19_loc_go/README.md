Dieses Programm dient dem Zählen von Zeilen.

Unit-Tests ausführen
```
go test
```

Der erste Test liefert als Ergebnis 0:
```
echo | ./loc 
```

Ein weiteres Beispiel liefert als Ergebnis: 2:
(-e parameter: enables interpretation of escape sequences)
```
echo -e "_\n_" | ./loc 
```
