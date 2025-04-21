# 🔍 Go Notify Server + Perspective AI

Este proyecto es un pequeño servidor escrito en **Golang** que permite enviar mensajes entre usuarios usando **WebSockets** y **HTTP REST**, con verificación de contenido moderado gracias a la API de **Perspective AI** de Google.

---

## ✨ Características

- ✅ WebSocket en tiempo real para recibir mensajes dirigidos.
- ✅ Endpoint HTTP para enviar mensajes a otros usuarios.
- ✅ Moderación automática de mensajes con [Perspective API](https://perspectiveapi.com/).
- ✅ Prevención de mensajes tóxicos, ofensivos o con discurso de odio.
- 🛡️ Basado en middleware limpio y simple para facilitar la extensión.

---

## 🧠 Requisitos

- Go 1.18 o superior
- Cuenta en [Google Cloud Console](https://console.cloud.google.com/) con acceso a Perspective API
- Clave API válida (ver abajo)

---

## ⚙️ Configuración

1. Crea una clave de API desde [Perspective API](https://www.perspectiveapi.com/#/start).
2. Exporta tu clave como variable de entorno:

```bash
export API_PERSPECTIVE_KEY=tu_clave_de_api
```

> ⚠️ Esta variable es **obligatoria**, sin ella la API no funcionará.

---

## 📦 Instalación y ejecución

```bash
git clone https://github.com/tu_usuario/notify-perspective
cd notify-perspective
go run main.go
```

---

## 🔌 Endpoints

### 📥 1. WebSocket: Recibir mensajes

`GET {{ _['url-base'] }}/notify?name={{name}}`

- Se conecta mediante WebSocket.
- El `name` representa el nombre del usuario que recibirá mensajes.
- El servidor enviará mensajes tipo `string` plano por el canal.

📌 Ejemplo:

```
ws://localhost:8080/notify?name=juan
```

---

### 📤 2. REST HTTP: Enviar mensajes

`POST {{ _['url-base'] }}/notify?name={{SENDER_NAME}}`

#### Body JSON:

```json
{
  "name": "juan",            // Receptor
  "message": "Hola amigo!"   // Mensaje
}
```

- El mensaje será evaluado por Perspective AI antes de ser entregado.
- Si el mensaje se considera **tóxico**, será reemplazado por `"*****"` antes de ser enviado al receptor.

---

## 🧪 Respuesta de errores (solo en caso de error interno)

```json
{
  "error": "Ocurrió un error interno al procesar el mensaje"
}
```

---

## 🛠️ Tecnologías usadas

- Go
- Gorilla WebSocket
- Perspective API (REST)
- Standard `net/http` + JSON

---

## 📜 Licencia

MIT License

---

## 🙌 Autor

Hecho con 💻 por [Juan David Sierra](https://github.com/jSierraB3991/)

