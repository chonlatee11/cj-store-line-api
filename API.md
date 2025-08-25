# API Documentation

## CJ Store LINE API

### Base URL
```
http://localhost:8080
```

### Health Check Endpoints

#### Liveness Probe
```http
GET /healthz
```

**Response:**
```json
{
  "status": "ok",
  "probe": "liveness"
}
```

#### Readiness Probe
```http
GET /readyz
```

**Response:**
```json
{
  "status": "ok",
  "probe": "readiness"
}
```

### Webhook Endpoints

#### Process Webhook Event
```http
POST /api/v1/webhook/line
Content-Type: application/json
```

**Request Body:**
```json
{
  "source": "ecommerce_system",
  "event_type": "message",
  "user_id": "U1234567890abcdef",
  "message": "Hello, this is a test message",
  "payload": {
    "order_id": "12345",
    "amount": 1500,
    "currency": "THB"
  }
}
```

**Response (Success):**
```json
{
  "status": "ok",
  "message": "webhook processed successfully",
  "correlationId": "uuid-correlation-id"
}
```

**Response (Error):**
```json
{
  "error": "failed to process webhook",
  "details": "error description",
  "correlationId": "uuid-correlation-id"
}
```

### Supported Event Types

1. **message** - Send a custom message to the user
2. **order_update** - Notify user about order status changes
3. **promotion** - Send promotional messages
4. **default** - Generic thank you message for unknown event types

### Request Headers

- `Content-Type: application/json` (required)
- `X-Correlation-ID: <uuid>` (optional, auto-generated if not provided)

### Error Codes

- `400 Bad Request` - Invalid request format or missing required fields
- `500 Internal Server Error` - Server error or failed to process webhook

### Example Usage

#### Send a Message Event
```bash
curl -X POST http://localhost:8080/api/v1/webhook/line \
  -H "Content-Type: application/json" \
  -H "X-Correlation-ID: 12345678-1234-1234-1234-123456789012" \
  -d '{
    "source": "store_system",
    "event_type": "message",
    "user_id": "U1234567890abcdef",
    "message": "Your order has been confirmed!",
    "payload": {
      "order_id": "ORD-2024-001",
      "total": 2500.00
    }
  }'
```

#### Send an Order Update Event
```bash
curl -X POST http://localhost:8080/api/v1/webhook/line \
  -H "Content-Type: application/json" \
  -d '{
    "source": "fulfillment_system",
    "event_type": "order_update",
    "user_id": "U1234567890abcdef",
    "payload": {
      "order_id": "ORD-2024-001",
      "status": "shipped",
      "tracking_number": "TH123456789"
    }
  }'
```
