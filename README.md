# HOC-Proyecto2
Problema del bombero con heurística de hormigas

Para ejecutar el programa es necesario tener golang instalado y colocar el proyecto dentro de la ruta $GOPATH/src/github/kaarla, una vez colocado en este punto, desde la carpeta raíz del proyecto ejecutar:

> HOC-Proyecto2 $ go install

Una vez hecho esto el proyecto se ejecuta con:

> HOC-Proyecto2 $ go run main.go [opcion]

El resultado de la ejecución arrojará código svg a la salida estándar, así que es necesario direccionar la salida a un archivo para visualizar las gráficas en el navegador.
La salida tendrá primero el dibujo del vecindario en el tiempo 0 y el siguiente el escenario final; en ocasiones se imprimirá este último más de una vez ya que más de una hormiga puede llegar a este resultado.

Para ejecutar el experimento con la gráfica de 50 vértices, con 3 incendios iniciales, ejecutar:

> HOC-Proyecto2 $ go run main.go 50

Se obtendrán los resultados:
>Semilla: 547
>Costo: 13.61111111111111
>Salvados: 37
>Bomberos usados: 14

Para ejecutar el experimento con la gráfica de 1000 vértices, con 3 incendios iniciales en las orillas, ejecutar:

> HOC-Proyecto2 $ go run main.go 1000-a

Se obtendrán los resultados:
>Semilla: 4845
>
>Costo: 12.594800000000001
>
>Salvados: 972
>
>Bomberos usados: 37

Para ejecutar el experimento con la gráfica de 1000 vértices, con 3 incendios iniciales en el centro, ejecutar:

> HOC-Proyecto2 $ go run main.go 1000-b

Se obtendrán los resultados:
>Semilla: 8710526160774049443
>
>Costo: 141.14880000000002
>
>Salvados: 937
>
>Bomberos usados: 78

Si no se elige opción, por omisión se ejecutará para la gráfica de 100 vértices y 4 incendios iniciales.
>Semilla: 54565
>
>Costo: 13.537500000000001
>
>Salvados: 81
>
>Bomberos usados: 19
