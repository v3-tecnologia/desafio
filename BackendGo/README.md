# BackendGo

## Descrição
API para receber dados de giroscópio, GPS e fotos. Armazena os dados em MongoDB e integra com AWS Rekognition para comparação de fotos.

## Endpoints

### POST /telemetry/gyroscope
Recebe dados de giroscópio:
```json
{
  "x": 0.123,
  "y": 0.456,
  "z": 0.789,
  "timestamp": 1633036800,
  "deviceId": "device-123"
}
