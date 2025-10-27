# 🚀 Patrones de Concurrencia en Go

> Repositorio de ejemplos y experimentos con patrones de concurrencia en Go: pipelines, fan-out/fan-in, worker pools, pub/sub, control dinámico de workers, supervisores, barrier, context-based cancellation y más.

**Autor:** Fernando Vélez M.

---

## 📚 Descripción

Este repositorio reúne ejemplos prácticos (cada uno en su carpeta) que ilustran patrones de concurrencia comunes en Go. El objetivo es aprender, probar y comparar enfoques para coordinar goroutines, manejar cancelaciones, evitar condiciones de carrera y diseñar sistemas concurrentes más robustos.

---

## ⚙️ Requisitos

- Go 1.20+ (o la versión estable más reciente)
- `git` para clonar (opcional)
- (Opcional) `golangci-lint` para linteo y `gofmt` para formato

---

## ▶️ Cómo ejecutar los ejemplos

Se asume que cada carpeta contiene un `main` ejecutable (por ejemplo `main.go` o `<pattern>.go`). Comandos típicos:
```bash
# clonar repo
git clone https://github.com/Fernandovm13/PatronesConcurrencia.git
cd PatronesConcurrencia

# ejecutar un ejemplo (desde la raíz)
go run ./ProCon           # Productor-Consumidor
go run ./Fan-out_Fan-in   # Fan-out / Fan-in
go run ./Pipeline         # Pipeline
go run ./Worker_pool      # Worker pool
go run ./Pub-Sub          # Publish-Subscribe
# ...y así con cada carpeta
```

---

## 🔍 Qué hace cada ejemplo

### `ProCon` — Productor / Consumidor 🛒

- **Propósito:** desacoplar el productor que genera trabajos y el/los consumidores que los procesan (cola intermedia por canal).
- **Uso esperado:** ejemplo con canal con buffer, productor que cierra el canal y consumidor(es) leyendo con `range`.
- **Qué probar:** cerrar el canal correctamente, variar el tamaño del buffer, ejecutar con `-race`.

### `Fan-out_Fan-in` — Paralelismo controlado ⚖️

- **Propósito:** distribuir trabajo a varios workers (fan-out) y recolectar resultados (fan-in).
- **Patrones clave:** `sync.WaitGroup` para coordinar cierre del canal de resultados.
- **Qué probar:** aumentar el número de workers y medir throughput/latencia.

### `Pipeline` — Procesamiento por etapas 🧩

- **Propósito:** componer etapas donde la salida de una alimenta la siguiente mediante canales.
- **Precaución:** propagar cancelaciones para evitar goroutines huérfanas.

### `Worker_pool` — Pool de trabajadores 🧰

- **Propósito:** limitar el número de goroutines concurrentes para controlar uso de recursos.
- **Consejo:** elegir pool size acorde a CPU/IO.

### `Future_Promise` — Resultado asíncrono 🔮

- **Propósito:** demostrar cómo retornar un valor "en el futuro" mediante canales (simular future/promise).
- **Recomendación:** incluir manejo de errores y timeouts.

### `Barrier` — Barrera de sincronización 🧱

- **Propósito:** esperar que N goroutines lleguen a un punto antes de continuar (ej. `sync.WaitGroup`).
- **Cuidado:** evitar bloqueos si alguna goroutine falla o tarda demasiado.

### `Pub-Sub` — Publicador / Suscriptores 📣

- **Propósito:** publicar eventos a varios suscriptores desacoplados.
- **Riesgo:** suscriptores lentos pueden bloquear al publicador — usar buffers o políticas de drop.

### `Select_Multiplexing` — `select` y timeouts ⏱️

- **Propósito:** mostrar multiplexación de canales, timeouts y prioridades.
- **Buen hábito:** en loops intensivos usar `time.NewTimer` para evitar fugas de timers.

### `Fan_Controller_Pattern` — Escalado dinámico ⚙️

- **Propósito:** ajustar dinámicamente el número de workers según la carga (longitud de la cola).
- **Complejidad:** scale-down seguro (parar workers) requiere mecanismos adicionales (quit channels, backoff, hysteresis).

### `Supervisor_Pattern` — Reinicio ante fallos 🛡️

- **Propósito:** supervisar goroutines, recuperar `panic` y reiniciar tareas con control (backoff).
- **Sugerencia:** implementar límites y reporting para evitar bucles de reinicio infinito.

