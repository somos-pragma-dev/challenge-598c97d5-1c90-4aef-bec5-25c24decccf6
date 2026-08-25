# Diseño de microservicios en Go

Estás encargado de diseñar un microservicio en Go para un sistema de procesamiento de pagos en una fintech. El microservicio debe manejar la solicitud de pagos, validar los datos de entrada, interactuar con un servicio externo para verificar la disponibilidad de fondos y registrar el resultado de la operación. Los pagos pueden provenir de diferentes canales (web, móvil, API) y deben ser idempotentes. El sistema debe manejar un volumen de 10 000 solicitudes por segundo en hora pico y garantizar una latencia máxima de 500ms. Debes decidir cómo estructurar el microservicio, cómo manejar los errores y cómo asegurar la idempotencia de las operaciones.

## Informacion General

| Campo | Valor |
|-------|-------|
| **Tema** | Microservicios en Go |
| **Nivel** | junior-l2 |
| **Tipo** | practical |
| **Tiempo estimado** | 8 horas |

## Fases del Reto

### Fase 0: Configuración del Proyecto

**Objetivo:** Obtener el proyecto base funcional enviando el Código Base a un asistente de IA, que lo analizará, corregirá errores y generará un ZIP listo para usar.

**Tiempo estimado:** 15-30 minutos

**Instrucciones:**

- Asegúrate de tener instalado para ejecutar el proyecto: Go 1.21+, VS Code o GoLand.
- Copia todo el contenido del campo **Código Base** de este reto — incluyendo el texto de instrucciones que aparece al inicio.
- Abre un asistente de IA (Claude en claude.ai, ChatGPT o Gemini — se recomienda Claude), pega el contenido copiado en el chat y envíalo.
- El asistente analizará los archivos, corregirá errores y generará un archivo ZIP descargable. Descárgalo y extráelo en la carpeta donde quieras trabajar.
- Ejecuta `go build ./...`. Si no hay errores, estás listo.

**Entregable:** El proyecto compila/arranca sin errores.

<details>
<summary>Pistas de conocimiento</summary>

- Copia el Código Base completo incluyendo el texto de instrucciones al inicio — esas instrucciones le indican al asistente exactamente qué hacer con los archivos.
- Si el asistente no genera el ZIP automáticamente al terminar el análisis, escríbele: "genera el ZIP ahora".
- Si el proyecto tiene errores al arrancar, comparte el mensaje de error con el mismo asistente para que lo corrija.

</details>

### Fase 1: Definición del microservicio

**Objetivo:** Definir la estructura y las responsabilidades del microservicio.

**Tiempo estimado:** 2 horas

**Instrucciones:**

- Identifica los canales de entrada y los datos que cada uno debe proporcionar.
- Define las validaciones necesarias para los datos de entrada.
- Especifica cómo el microservicio interactuará con el servicio externo para verificar la disponibilidad de fondos.

**Entregable:** Diagrama de relaciones que muestra los componentes del microservicio y sus interacciones.

<details>
<summary>Pistas de conocimiento</summary>

- Considera cómo los diferentes canales pueden afectar la estructura del microservicio.
- Piensa en las validaciones que pueden fallar y cómo manejar esos casos.

</details>

### Fase 2: Implementación de la idempotencia

**Objetivo:** Implementar la idempotencia en las operaciones del microservicio.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Diseña un mecanismo para asegurar que las operaciones sean idempotentes.
- Especifica cómo manejarás los casos donde la idempotencia falla.
- Define los umbrales de latencia y volumen que el microservicio debe cumplir.

**Entregable:** Descripción del mecanismo de idempotencia y cómo se integra con el microservicio.

<details>
<summary>Pistas de conocimiento</summary>

- Considera el uso de claves de idempotencia y cómo almacenarlas.
- Piensa en los modos de falla que pueden afectar la idempotencia y cómo mitigarlos.

</details>

### Fase 3: Integración con servicio externo

**Objetivo:** Integrar el microservicio con el servicio externo para verificar la disponibilidad de fondos.

**Tiempo estimado:** 3 horas

**Instrucciones:**

- Especifica cómo el microservicio se comunicará con el servicio externo.
- Define los casos de error que pueden ocurrir y cómo manejarlos.
- Asegura que el microservicio cumpla con los umbrales de latencia y volumen definidos.

**Entregable:** Descripción de la integración con el servicio externo y cómo se manejan los errores.

<details>
<summary>Pistas de conocimiento</summary>

- Considera el uso de protocolos de comunicación eficientes.
- Piensa en los modos de falla del servicio externo y cómo mitigarlos.

</details>

## Dimensiones Evaluadas

- **queEs**: ¿Qué es un microservicio y por qué se usa en este contexto?
- **paraQueSirve**: ¿Para qué sirve la idempotencia en este microservicio?
- **comoSeUsa**: ¿Cómo se usa la idempotencia en las operaciones del microservicio?
- **erroresComunes**: ¿Qué errores comunes pueden ocurrir al integrar con el servicio externo y cómo se manejan?
- **queDecisionesImplica**: ¿Qué decisiones implica la implementación de la idempotencia y la integración con el servicio externo?

## Criterios de Evaluacion

- Definición clara de la estructura y responsabilidades del microservicio.
- Mecanismo de idempotencia implementado y explicado.
- Integración con el servicio externo especificada y errores comunes manejados.
- Cumplimiento de los umbrales de latencia y volumen definidos.

---

*Reto generado automaticamente por Challenge Generator - Pragma*
