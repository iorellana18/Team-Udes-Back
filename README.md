# E-Hackaton Backend
API REST implementada en Golang de la aplicación E-Hackaton

## Estructura de los directorios
- `auth/`: Autentificación para el ingreso a la API REST.
- `config/`: Archivo de configuración para la ejecución del programa.
- `database`: Script de la base de datos del proyecto.
- `db/`: Conexión a la base de datos.
- `migrate`: Migra los datos de la base de datos al índice invertido.
- `models/`: Estructuras de cada una de las tablas de la base de datos con su respectivo CRUD.
- `routes/`: Rutas e implementación de servicios de la API REST.
- `utils/`: Misceláneo []() de utilidades para el funcionamiento del programa.

## Compilación del proyecto
Para actualizar las dependencias del proyecto, es necesario tener instalado `glide`, por lo que posterior a esto, debe ejecutarse `glide up` y luego compilarlo con `go build`.

## Ejecución del proyecto
Para su ejecución, es necesaria la carpeta `config` en el mismo directorio donde se va a ejecutar el programa compilado.
